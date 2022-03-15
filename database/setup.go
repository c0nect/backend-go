package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // create new db connection

	if err != nil {
		panic("error connecting to database...")
	}

	RunMigrations(db) // run migrations
	DB = db // receive instance of db
}

func GetDatabase() *gorm.DB {
	return DB
}