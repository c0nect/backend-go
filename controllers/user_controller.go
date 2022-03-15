package controllers

import (
	"MoviesBack/database"
	"MoviesBack/models"
	"MoviesBack/services"
	"log"

	"github.com/gin-gonic/gin"
)

const (
	USER  = "user"
	ADMIN = "admin"
)

func FindUsers(c *gin.Context) {
	var users []models.User

	err := database.GetDatabase().Find(&users).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "can't find users" + err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}

func CreateUser(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	// println(user.Password)
	// println(user.Role)
	result := database.GetDatabase().Where("email = ?", user.Email).Find(&user)
	if result.RowsAffected > 0 { // if user with email exist
		c.JSON(400, gin.H{
			"error": "user with this email already exists",
		})

		// println(user.Email)
		return
	}

	switch user.Role {
	case USER:
		user.Role = USER
	case ADMIN:
		user.Role = USER
		log.Println("############################################################")
		log.Println("[WARN] user role is admin, but it's not allowed")
		log.Println("you need set role admin directly in database")
		log.Println("############################################################")

		c.JSON(400, gin.H{
			"error": "user role is admin, but it's not allowed, try user",
		})

		return
	default:
		user.Role = USER
	}

	user.Password = services.SHA256Encrypt(user.Password) // encrypt password
	err = database.GetDatabase().Create(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "can't create user" + err.Error(),
		})

		return
	}

	// finnaly create a user
	c.JSON(200, gin.H{
		"message": "user created",
	})
}

// find user by id
func FindUser(c *gin.Context) {
	var user models.User

	err := database.GetDatabase().Where("id = ?", c.Param("id")).First(&user).Error // find user by email
	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})

		return // return if user not found
	}

	// send json response
	c.JSON(200, gin.H{
		"user": user,
	})
}

func FindUserByEmail(c *gin.Context) {
	var user models.User

	err := database.GetDatabase().Where("email = ?", c.Param("email")).First(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})

		return
	}

	c.JSON(200, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		// "role":      user.Role,
		"image_url": user.ImageUrl,
	})
}

func DeleteUser(c *gin.Context) {
	var user models.User

	err := database.GetDatabase().Where("id = ?", c.Param("id")).First(&user).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})

		return // return if user not found
	}

	err = database.GetDatabase().Delete(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "can't delete user" + err.Error(),
		})

		return
	}

	// send json response
	c.Status(204)
}
