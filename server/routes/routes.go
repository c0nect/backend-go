package routes

import (
	"MoviesBack/controllers"
	"MoviesBack/server/middlewares"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(r *gin.Engine) *gin.Engine {
	main := r.Group("/api/v1")
	{
		user := main.Group("/user")
		{
			user.GET("/:email", controllers.FindUserByEmail)
		}
		users := main.Group("/users")
		{
			users.GET("/", controllers.FindUsers)
			users.POST("/", controllers.CreateUser)
			users.GET("/:id", controllers.FindUser)
			users.DELETE("/:id", controllers.DeleteUser)
		}
		movies := main.Group("/movies", middlewares.Auth())
		{
			movies.GET("/", controllers.FindMovies)        // all movies
			movies.POST("/", controllers.CreateMovie)      // new movie
			movies.GET("/:id", controllers.FindMovie)      // movie by id
			movies.PATCH("/:id", controllers.UpdateMovie)  // update movie
			movies.DELETE("/:id", controllers.DeleteMovie) // delete movie
		}

		login := main.Group("/login")
		{
			login.POST("/", controllers.Login)
		}
	}

	return r
}
