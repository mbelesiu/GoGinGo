package database

import (
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

func NewDBOptions() *pg.Options {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	log.Println("Hello Matthew: ", dbPassword)
	return &pg.Options{
		// Addr:     "172.23.0.2:5432",
		Database: "rgb",
		User:     "postgres",
		Password: dbPassword,
	}
}
