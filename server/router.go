package server

import (
	"fmt"
	"time"

	"github.com/Piyushhbhutoria/go-gin-boilerplate/controllers"
	"github.com/Piyushhbhutoria/go-gin-boilerplate/docs"

	// "github.com/Piyushhbhutoria/go-gin-boilerplate/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST"}
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.MaxAge = 1 * time.Minute
	router.Use(cors.New(corsConfig))

	// router.Use(middlewares.AuthMiddleware())

	health := new(controllers.HealthController)
	user := new(controllers.UserController)

	router.GET("/health", health.Status)

	// User routes
	router.GET("/users", user.GetUsers)
	router.POST("/users", user.CreateUser)
	router.GET("/users/:id", user.GetUser)

	// Swagger documentation
	docs.SwaggerInfo.Title = "Go Gin Boilerplate API"
	docs.SwaggerInfo.Description = "A RESTful API built with Go, Gin, and GORM"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router

}
