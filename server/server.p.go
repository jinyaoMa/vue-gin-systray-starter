package server

import (
	"app/server/routes"
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
)

var server *Server

func (s *Server) prepare(swag bool) {
	handler := gin.Default()
	port := fmt.Sprintf(":%d", s.config.Port)
	portTls := fmt.Sprintf(":%d", s.config.PortTls)
	manager := &autocert.Manager{
		Prompt: autocert.AcceptTOS,
		Cache:  autocert.DirCache(s.config.CertDir),
	}
	tlsConfig := manager.TLSConfig()

	routes.Init(s.config, handler, swag)

	s.http = &http.Server{
		Addr: port,
		Handler: manager.HTTPHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			target := "https://" + strings.Replace(r.Host, port, portTls, 1) + r.RequestURI
			http.Redirect(w, r, target, http.StatusMovedPermanently)
		})),
	}

	tlsConfig.GetCertificate = s.getSelfSignedOrLetsEncryptCert(manager)
	s.https = &http.Server{
		Addr:      portTls,
		TLSConfig: tlsConfig,
		Handler:   handler,
	}
}

func (s *Server) startRedirecter() {
	err := s.http.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		s.logger.Fatalf("Redirecter listen: %v\n", err)
		return
	}
}

func (s *Server) stopRedirecter() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = s.http.Shutdown(ctx)
	if err != nil {
		s.logger.Fatalf("Redirecter with port %d shutdown: %v\n", s.config.Port, err)
		return
	}

	s.logger.Printf("Redirecter with port %d is exiting!\n", s.config.Port)
	return
}

func (s *Server) startServer() {
	err := s.https.ListenAndServeTLS("", "")
	if err != nil && err != http.ErrServerClosed {
		s.logger.Fatalf("Server (TLS) listen: %v\n", err)
		return
	}
}

func (s *Server) stopServer() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = s.https.Shutdown(ctx)
	if err != nil {
		s.logger.Fatalf("Server (TLS) with port %d shutdown: %v\n", s.config.PortTls, err)
		return
	}

	s.logger.Printf("Server (TLS) with port %d is exiting!\n", s.config.PortTls)
	return
}

func (s *Server) resetGin(logger *log.Logger) {
	gin.DefaultWriter = logger.Writer()
	if s.isDev {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}
}

func (s *Server) getSelfSignedOrLetsEncryptCert(certManager *autocert.Manager) func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	return func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		dirCache, ok := certManager.Cache.(autocert.DirCache)
		if !ok {
			dirCache = "certs"
		}

		keyFile := filepath.Join(string(dirCache), hello.ServerName+".key")
		crtFile := filepath.Join(string(dirCache), hello.ServerName+".crt")
		certificate, err := tls.LoadX509KeyPair(crtFile, keyFile)
		if err != nil {
			s.logger.Printf("%s\nFalling back to Letsencrypt\n", err)
			return certManager.GetCertificate(hello)
		}
		s.logger.Println("Loaded selfsigned certificate.")
		return &certificate, err
	}
}
