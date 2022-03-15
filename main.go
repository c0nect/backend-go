package main

import (
	"MoviesBack/database"
	"MoviesBack/server"
)

func main() {
	database.ConnectDB()
	server := server.NewServer()
	server.Run()
}

// lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec euismod,
