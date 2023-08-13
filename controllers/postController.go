package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Userasma/Rihal-Project/initializers"
	"github.com/Userasma/Rihal-Project/models"
	"github.com/gin-gonic/gin"
)

// type UserForm struct {
// 	FirstName   string    `json:"first_name"`
// 	LastName    string    `json:"last_name"`
// 	Email       string    `json:"email"`
// 	DateOfBirth time.Time `json:"date_of_birth"`
// }

// type UserUpdateForm struct {
// 	FirstName   string    `json:"first_name"`
// 	LastName    string    `json:"last_name"`
// 	Email       string    `json:"email"`
// 	DateOfBirth time.Time `json:"date_of_birth"`
// }

func UserCreate(c *gin.Context) {
	var userForm UserForm
	if err := c.ShouldBindJSON(&userForm); err != nil {
		//fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//fmt.Println("User form:", userForm)

	user := models.User{
		FirstName:   userForm.FirstName,
		LastName:    userForm.LastName,
		Email:       userForm.Email,
		DateOfBirth: userForm.DateOfBirth,
	}

	//This to create user database
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		fmt.Println("Failed to create user:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	//fmt.Println("User created successfully:", user)
	//c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	c.Redirect(http.StatusSeeOther, "/user/list")
}

func CalculateAge(dateOfBirth time.Time) int {
	now := time.Now()
	age := now.Year() - dateOfBirth.Year()

	if now.YearDay() < dateOfBirth.YearDay() {
		age--
	}
	return age
}

// func CreatePost(c *gin.Context) {
//     var postForm UserForm
//     if err := c.ShouldBindJSON(&postForm); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }

//	    c.JSON(http.StatusOK, gin.H{"message": "Post created successfully"})
//	}
func GetUsers() []models.User {
	// Retrieve the users from the database
	var users []models.User
	initializers.DB.Find(&users)
	return users
}
func GetUserByID(userID string) *models.User {
	// Retrieve the user from the database based on the provided ID
	var user models.User
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
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found ;)"})
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
	fmt.Println("Deleting user with ID:", id)
	// Delete the user
	result := initializers.DB.Delete(&models.User{}, id)
	if result.Error != nil {
		fmt.Println("Failed to delete user:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	fmt.Println("User deleted successfully:", id)

	// respond
	//c.Redirect(http.StatusFound, "/users/list")
}
func ShowUserForm(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", nil)
}

func ShowUserList(c *gin.Context) {
	users := GetUsers()

	c.HTML(http.StatusOK, "userList.html", gin.H{
		"users": users,
	})
	//fmt.Println("User deleted successfully:", users)
}

func ShowUpdateUserForm(c *gin.Context) {
	userID := c.Param("id")
	user := GetUserByID(userID)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.HTML(http.StatusOK, "updateUser.html", gin.H{
		"user": user,
	})
}

// respond
//c.Redirect(http.StatusFound, "/users/list")
