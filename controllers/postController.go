package controllers

import (
	//"net/http"
	"net/http"
	"time"

	"github.com/Userasma/Rihal-Project/initializers"
	"github.com/Userasma/Rihal-Project/models"
	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {
	// Bind JSON data to the UserForm struct
	var userForm UserForm
	if err := c.ShouldBindJSON(&userForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new user based on the form data
	user := models.Main{
		FirstName:   userForm.FirstName,
		LastName:    userForm.LastName,
		Email:       userForm.Email,
		DateOfBirth: userForm.DateOfBirth,
	}

	// Create the user in the database
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Send a success response
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// Function to calculate age based on the date of birth
func CalculateAge(dateOfBirth time.Time) int {
	now := time.Now()
	age := now.Year() - dateOfBirth.Year()

	//If the current day of the year is less than the user's day of the year,
	//  we decrement the age.
	if now.YearDay() < dateOfBirth.YearDay() {
		age--
	}

	return age
}

func GetUsers() []models.Main {
	// Retrieve the users from the database
	var users []models.Main
	initializers.DB.Find(&users)
	return users
}

func GetUserByID(userID string) *models.Main {
	// Retrieve the user from the database based on the provided ID
	var user models.Main
	if err := initializers.DB.First(&user, userID).Error; err != nil {
		// Handle the error if the user is not found or any other database error occurs
		return nil
	}
	return &user
}

func UserUpdate(c *gin.Context) {
	//get id of url
	ID := c.Param("id")
	// get the user by the given id from the url
	user := GetUserByID(ID)
	if user == nil {
		// Handle the case where the user is not found
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Parse and bind the form data to the UserUpdateForm struct
	var form UserUpdateForm
	if err := c.ShouldBind(&form); err != nil {
		// Handle the error if form data validation fails
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the user with the new details
	user.FirstName = form.FirstName
	user.LastName = form.LastName
	user.Email = form.Email
	user.DateOfBirth = form.DateOfBirth

	// Save the updated user to the database
	if err := initializers.DB.Save(user).Error; err != nil {
		// Handle the error if the user update fails
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// Redirect to the user details page
	//c.Redirect(http.StatusSeeOther, "/users/list")
}

func UserDelete(c *gin.Context) {
	// get id off url
	id := c.Param("id")
	// Delete the posts
	initializers.DB.Delete(&models.Main{}, id)
	// respond
	//c.Redirect(http.StatusFound, "/users/list")
}
