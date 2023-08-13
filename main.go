package main

import (
	"github.com/Userasma/Rihal-Project/initializers"
	"github.com/Userasma/Rihal-Project/routers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := routers.SetupRouter()

	r.Run()

}

// gin.SetMode(gin.DebugMode)
// r := gin.Default()

// r.Static("/static", "./static") // Serve static files like CSS and JS
// r.LoadHTMLGlob("templates/*")    // Load HTML templates from the 'templates' directory

// r.GET("/users/form", controllers.ShowUserForm)
// r.GET("/users/list", controllers.ShowUserList)
// r.GET("/users/:id/update", controllers.ShowUpdateUserForm)

// r.POST("/users", controllers.UserCreate)       // Create a user
// r.PUT("/users/:id", controllers.UserUpdate)    // Update a user by ID
// r.DELETE("/users/:id", controllers.UserDelete) // Delete a user by ID

// r.Run(":3000")
// func main() {
// 	gin.SetMode(gin.DebugMode)
// 	r := gin.Default()

// 	r.POST("/users", controllers.UserCreate)       // Create a user
// 	r.PUT("/users/:id", controllers.UserUpdate)    // Update a user by ID
// 	r.DELETE("/users/:id", controllers.UserDelete) // Delete a user by ID

// 	r.Run(":3000")
// 	//seed.SeedUers(initializers.DB)
// 	//r:=routers.SetupRouter()
// 	// r.Run()
// }

// func main() {
// 	gin.SetMode(gin.DebugMode)
// 	r := gin.Default()
// 	//r.GET("/form", controllers.ShowUserForm)
// 	//r.GET("/list", controllers.ShowUserList)
// 	r.POST("/posts", controllers.UserCreate)
// 	r.POST("/", controllers.UserCreate)
// 	r.GET("/:id", controllers.UserUpdate)
// 	r.POST("/:id/delete", controllers.UserDelete)
// 	r.Run(":3000")
// }

// func main() {
// 	seed.SeedUsers(initializers.DB)
// 	r := routers.SetupRouter()

// 	// Run the server
// 	r.Run()
// }
