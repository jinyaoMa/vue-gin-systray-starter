package routes

import (
	"app/config"
	"app/database"
	"app/server/models"
	"app/server/routes/users"
	_ "app/server/swagger"
	"app/utils/logger"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title app
// @version 0.0.0
// @description "app"

// @contact.name Github Issues
// @contact.url https://github.com/jinyaoMa/vue-gin-systray-starter/issues

// @license.name MIT
// @license.url https://github.com/jinyaoMa/vue-gin-systray-starter/blob/main/LICENSE

// @BasePath /api

// @securityDefinitions.apikey BearerToken
// @in header
// @name Authorization

func Init(config *config.Server, handler *gin.Engine, swag bool) {
	logger, ok := logger.GetInstance()
	if ok {
		db := database.Connect(logger.Database, config.Database)
		db.AutoMigrate(&models.User{})
	}

	if gin.Mode() == gin.DebugMode {
		handler.Use(cors.New(cors.Config{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET", "POST"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
			AllowCredentials: false,
			MaxAge:           12 * time.Hour,
		}))
	} else {
		port := fmt.Sprintf(":%d", config.Port)
		portTls := fmt.Sprintf(":%d", config.PortTls)
		handler.Use(cors.New(cors.Config{
			AllowOrigins: []string{
				"http://" + config.Origin + ":" + port,
				"https://" + config.Origin + ":" + portTls,
			},
			AllowMethods:     []string{"GET", "POST"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))

		path, err := os.Executable()
		if err == nil {
			handler.Static("/assets", filepath.Join(filepath.Dir(path), "./www/assets"))
			handler.StaticFile("/favicon.ico", filepath.Join(filepath.Dir(path), "./www/favicon.ico"))
			handler.StaticFile("/", filepath.Join(filepath.Dir(path), "./www/index.html"))
			handler.GET("/index.html", func(c *gin.Context) {
				c.Redirect(http.StatusFound, "/")
			})
		}
	}

	api := handler.Group("/api")

	_users := api.Group("/users")
	{
		_users.GET("/all", users.All())
	}

	if swag {
		handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
