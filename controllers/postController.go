package controllers

import (
	"github.com/Userasma/Rihal-Project/initializers"
	"github.com/Userasma/Rihal-Project/models"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var main models.Main
	if err := c.ShouldBindJSON(&main); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := initializers.DB.Create(&main)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{
		"main": main,
	})
}

func Index(c *gin.Context) {
	var mains []models.Main
	initializers.DB.Find(&mains)

	c.JSON(200, gin.H{
		"mains": mains,
	})
}

func Show(c *gin.Context) {
	id := c.Param("id")

	var main models.Main
	result := initializers.DB.First(&main, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(200, gin.H{
		"main": main,
	})
}

func Update(c *gin.Context) {
	id := c.Param("id")

	var main models.Main
	result := initializers.DB.First(&main, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}

	var updatedMain models.Main
	if err := c.ShouldBindJSON(&updatedMain); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Model(&main).Updates(updatedMain)

	c.JSON(200, gin.H{
		"main": main,
	})
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	var main models.Main
	result := initializers.DB.First(&main, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}

	initializers.DB.Delete(&main)

	c.Status(200)
}
