package controllers

import (
	//"html/termplete"
	//"net/http"
	//"strings"
	"time"
	//"github.com/gin-goinc/gin"
)

type UserForm struct{
	FirstName string `form:"FirstName" binding:"required"`
	LastName  string `form:"LastName" binding:"required"`
	Email     string `form:"Email" binding:"required,email"`
	DateOfBirth time.Time `form:"DateOfBirth" binding:"required" time_format:"2006-01-02"`
}

type UserUpdateForm struct{
	FirstName string `form:"FirstName" binding:"required"`
	LastName  string `form:"LastName" binding:"required"`
	Email     string `form:"Email" binding:"required,email"`
	DateOfBirth time.Time `form:"DateOfBirth" binding:"required" time_format:"2006-01-02"`
}

// func ShowUserForm(c *gin.Context){
// 	formHTML := RenderFormHTML()
// 	c.HTML(http.StatusOK, "base.html", gin.H{
// 		"content": template.HTML(formHTML),
// 	})

// }

// func ShowUserList(c *gin.Context) {
// 	users := GetUser()
// 	totalUser := len(users)

// 	var totalAge
// }