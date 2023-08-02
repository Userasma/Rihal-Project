package main

import (
	"github.com/LENOVO/student-statistics-app/Rihal-Project/initializers"
	"github.com/LENOVO/student-statistics-app/Rihal-Project/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})

}
