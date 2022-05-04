package main

import (
	"OSProject1/src/controllers"
	"OSProject1/src/traffic"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.LoadHTMLFiles(
		"src/views/index.html",
		"src/views/newPost.html",
		"src/views/naming.html",
		"src/views/traffic.html",
	)
	r.GET("/", func(ctx *gin.Context) { ctx.Redirect(http.StatusFound, "/posts") })

	// Post
	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/posts/create", controllers.GetCreatePost)
	r.POST("/posts/create", controllers.CreatePost)

	// naming
	r.GET("/naming", controllers.GetNaming)
	r.POST("/naming", controllers.Naming)

	// traffic simulator
	r.GET("/traffic", traffic.Get)
	r.GET("/traffic/start", traffic.Start)
	r.GET("/traffic/end", traffic.End)
	r.GET("/traffic/addCar", traffic.AddCar)
	r.GET("/traffic/display", traffic.Display)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
