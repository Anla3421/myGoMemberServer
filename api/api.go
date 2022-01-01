package api

import (
	"server/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Api() {
	docs.SwaggerInfo.Title = "Jared's Simple Member Server API by Golang"
	docs.SwaggerInfo.Description = "由golang撰寫的簡單會員服務API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "127.0.0.1:9000"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()

	// r.Use(cors.Default())
	r.Use(cors.New(CorsConfig()))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	Router(r)

	listenPort := "9000"
	r.Run(":" + listenPort)
}

func CorsConfig() cors.Config {
	corsConf := cors.DefaultConfig()
	corsConf.AllowOriginFunc = func(origin string) bool {
		return true
	}

	corsConf.AllowCredentials = true
	corsConf.AllowBrowserExtensions = true
	corsConf.AllowWildcard = true

	// corsConf.AllowAllOrigins = true
	corsConf.AllowMethods = []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"}
	corsConf.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin",
		"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers",
		"Methods"}
	return corsConf
}
