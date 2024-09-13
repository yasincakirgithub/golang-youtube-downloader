package router

import (
	"GOLANG/api/controllers"
	"GOLANG/api/middlewares"

	"github.com/gin-gonic/gin"
)

func GetRoute(r *gin.Engine) {

	r.LoadHTMLGlob("api/templates/*")
	// homepage
	r.GET("/", func(c *gin.Context) {
		// index.html dosyasını render eder
		c.HTML(200, "homepage.html", gin.H{
			"title": "Ana Sayfa",
		})
	})

	// User Route
	api_users := r.Group("/api/users")
	api_users.POST("/register", controllers.Register)
	api_users.POST("/auth", controllers.Auth)
	api_users.GET("/me", middlewares.CheckAuth, controllers.GetUserProfile)

	//
	api_youtube_downloader := r.Group("/api/youtube")
	api_youtube_downloader.POST("/download", controllers.DownloadYoutubeVideo)
}
