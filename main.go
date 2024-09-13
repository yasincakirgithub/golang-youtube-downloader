package main

import (
	"GOLANG/api/initializers"
	r "GOLANG/api/router"
	docs "GOLANG/docs"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

// @title           Simple JWT Auth API With Gin Framework
// @version         1.0
// @contact.name   Yasin Çakır
// @contact.url    https://www.linkedin.com/in/yasincakir26/
// @host      127.0.0.1:3000
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	// Gin Framework
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GetRoute(router)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Static("/static", "./static")
	router.Run(":" + os.Getenv("PORT"))

}
