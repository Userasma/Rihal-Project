package main

import (
	"github.com/Userasma/Rihal-Project/initializers"
	"github.com/Userasma/Rihal-Project/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Main{})

}
