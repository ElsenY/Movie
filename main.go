package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"test-msbu/core"
	"test-msbu/handler"
	"test-msbu/routes"
	"test-msbu/services"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func initDBConn() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := initDBConn()

	// manual dependency injection, can use other library like facebookgo/inject for future development
	movieCore := core.NewMovieCore(db)
	movieServices := services.NewMovieServices(movieCore)
	MovieHandler := handler.NewMovieHandler(movieServices)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	r := routes.SetupRouter(MovieHandler)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8081")
}
