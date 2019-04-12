package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/mammenj/ginusers/users/db"
	"github.com/mammenj/ginusers/users/models"
	"github.com/mammenj/ginusers/users/security"

	"log"
	"net/http"
)

//GetUsers gets all users in the table
func GetUsers(c *gin.Context) {
	var users []models.User
	log.Println("GetUsers from db")
	db := db.GetDB()
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

//GetUser get the User by id
func GetUser(c *gin.Context) {
	log.Println("GetUser from db")
	var user models.User
	id := c.Param("id")
	db := db.GetDB()
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("Failed to GetUser in db")
		return
	}
	c.JSON(http.StatusOK, user)
}

//CreateUser creates the user by User object
func CreateUser(c *gin.Context) {
	log.Println("CreateUser in db")
	var user models.User
	var db = db.GetDB()
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println("Failed to create user in db")
		return
	}
	// hash the password
	user.Password = security.HashAndSalt([]byte(user.Password))

	db.Create(&user)
	c.JSON(http.StatusOK, &user)
}

//UpdateUser updates the user by User object provided
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	log.Printf("UpdateUser in db %v", id)
	var user models.User

	db := db.GetDB()
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("Failed to UpdateUser in db")
		return
	}
	c.BindJSON(&user)
	db.Save(&user)
	c.JSON(http.StatusOK, &user)
}

//DeleteUser deletes user by the userid given
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	log.Printf("DeleteUser in db %v", id)
	var user models.User
	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println("Failed to DeleteUser in db")
		return
	}

	db.Delete(&user)
}
