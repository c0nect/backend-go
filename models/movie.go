package models

// my types

type Movie struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PosterPath  string `json:"poster_path"`
	// BackdropPath string `json:"backdrop_path"`
}

type CreateMovieInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	PosterPath  string `json:"poster_path" binding:"required"`
}

// equals /\ but not required
type UpdateMovieInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PosterPath  string `json:"poster_path"`
}
