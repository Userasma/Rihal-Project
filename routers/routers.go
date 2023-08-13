package routers

import (
	"github.com/Userasma/Rihal-Project/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//to load HTML templates
	r.LoadHTMLGlob("templates/*.html")

	//Group routers under /users

	users := r.Group("/users")
	{
		users.GET("/form", controllers.ShowUserForm)
		users.GET("/list", controllers.ShowUserList)
		users.POST("/", controllers.UserCreate)
		//users.GET("/:id", controllers.ShowUserUpdate)
		users.POST("/:id/update", controllers.UserUpdate)
		users.POST("/:id/delete", controllers.UserDelete)
	}
	return r
}
