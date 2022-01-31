package routes

import (
	_ "app/swagger"
	"net/http"

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

func Init(r *gin.Engine, swag bool) {
	api := r.Group("/api")

	api.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	if swag {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
