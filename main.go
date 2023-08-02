package main

import (
	"github.com/Userasma/Rihal-Project/controllers"
	"github.com/Userasma/Rihal-Project/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB() //to esatblish the connection

}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostCreate)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.GET("/posts", controllers.PostIndex)
	//this work as a dynamic router
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("/posts/:id", controllers.PostDelete)
	//commit one

	r.Run(":3000")
}
