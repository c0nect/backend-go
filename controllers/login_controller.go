package controllers

import (
	"MoviesBack/database"
	"MoviesBack/models"
	"MoviesBack/services"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	db := database.GetDatabase()

	var login models.Login
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}
	// can bind json to struct

	var user models.User
	dbError := db.Where("email = ?", login.Email).First(&user).Error
	if dbError != nil {
		c.JSON(400, gin.H{
			"error": "can't find user",
		})

		return
	}
	// now i have a user, now i need to compare the password
	if user.Password != services.SHA256Encrypt(login.Password) {
		c.JSON(400, gin.H{
			"error": "wrong password",
		})

		return
	}

	token, err := services.NewJWTService().GenerateToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})

		return
	}

	// somente adicionar
	// cookies no browser
	// insomnia funciona
	// maybe? => https://stackoverflow.com/questions/63860373/how-to-use-cookie-inside-getserversideprops-method-in-next-js

	c.SetCookie("jwt", token, 60*60*2, "/", "", true, true) // 2 hours secure true, httponly true
	// c.SetCookie("jwt", token, 10, "", "localhost", true, false) // secure true, httonly false
	// println(token)
	c.JSON(200, gin.H{
		"token": token,
		"user":  user,
	})

}
