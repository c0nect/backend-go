package controllers

import (
	"MoviesBack/database"
	"MoviesBack/models"

	"github.com/gin-gonic/gin"
)

// GET /api/v1/movies
// Get all movies
func FindMovies(c *gin.Context) {
	var movies []models.Movie
	database.GetDatabase().Find(&movies)

	c.JSON(200, gin.H{
		"movies": movies,
	})
}

func CreateMovie(c *gin.Context) {
	// validate input
	var input models.CreateMovieInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	// movies info
	movie := models.Movie{
		Title:       input.Title,
		Description: input.Description,
		PosterPath:  input.PosterPath,
	}

	err = database.GetDatabase().Create(&movie).Error // create new movie
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	database.GetDatabase().Create(&movie)
		      
	// send json response
	c.JSON(200, gin.H{
		"movie": movie,
	})
}

func FindMovie(c *gin.Context) {
	// validate input
	var movie models.Movie

	err := database.GetDatabase().Where("id = ?", c.Param("id")).First(&movie).Error // find movie by id

	if err != nil {
		c.JSON(400, gin.H{
			"error": "movie not found",
		})

		return // return if movie not found
	}

	// send json response
	c.JSON(200, gin.H{
		"movie": movie,
	})
}

func UpdateMovie(c *gin.Context) {
	// Get model if exist
	var movie models.Movie
	if err := database.GetDatabase().Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(400, gin.H{"error": "Record not found!"})

		return
	}

	// Validate input
	var input models.UpdateMovieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}

	database.GetDatabase().Model(&movie).Updates(models.Movie{
		Title:       input.Title,
		Description: input.Description,
		PosterPath:  input.PosterPath,
	})

	c.JSON(200, gin.H{
		"movie": movie,
	})
}

func DeleteMovie(c *gin.Context) {
	var movie models.Movie

	err := database.GetDatabase().Where("id = ?", c.Param("id")).First(&movie).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "movie not found",
		})

		return
	}

	err = database.GetDatabase().Delete(&movie).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "can't delete movie" + err.Error(),
		})

		return
	}

	c.Status(204)
}
