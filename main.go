package main

import (
	"github.com/Userasma/Rihal-Project/controllers"
	"github.com/Userasma/Rihal-Project/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB() //to establish the connection

}

func main() {
	r := gin.Default()

	// Serve static files from the "frontend" folder
	r.Static("/static", "./frontend")

	r.POST("/posts", controllers.PostCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.GET("/posts", controllers.PostIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("/posts/:id", controllers.PostDelete)

	r.Run(":3000")
}
