package server

import (
	"app/config"
	"log"
	"net/http"
)

type Server struct {
	isDev     bool
	isRunning bool
	logger    *log.Logger
	config    *config.Server
	http      *http.Server // redirecter
	https     *http.Server // server (tls)
}

func GetInstance() (*Server, bool) {
	return server, server != nil
}

func New(logger *log.Logger, config *config.Server, isDev bool) *Server {
	if server != nil {
		if !server.isRunning {
			server.logger = logger
			server.config = config
			server.isDev = isDev
			server.resetGin(logger)
		}
		return server
	}

	server = &Server{
		logger: logger,
		config: config,
		isDev:  isDev,
	}
	server.resetGin(logger)
	return server
}

func (s *Server) Start(swag bool, loop bool) (ok bool) {
	if s.isRunning {
		return false
	}

	s.isRunning = true
	s.prepare(swag)
	if loop {
		go s.startRedirecter()
		s.startServer()
	} else {
		go s.startRedirecter()
		go s.startServer()
	}

	return true
}

func (s *Server) Stop() (ok bool) {
	if !s.isRunning {
		return false
	}

	s.isRunning = false
	s.stopRedirecter()
	s.stopServer()
	return true
}
