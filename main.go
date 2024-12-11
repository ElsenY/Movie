package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"test-msbu/core"
	"test-msbu/handler"
	"test-msbu/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db = make(map[string]string)

func setupRouter(mh handler.IMovieHandler) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// r.POST("/movie",mh.CreateMovieHandler)

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))

	admin := r.Group("/", gin.BasicAuth(gin.Accounts{
		os.Getenv("ADMIN_USERNAME"):  os.Getenv("ADMIN_PASSWORD"), // user:foo password:bar
	}))

	admin.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	admin.POST("movie",mh.CreateMovie)
	admin.PUT("movie/:id",mh.UpdateMovieById)
	admin.GET("mostviewedmovie",mh.GetMostViewedMovie)

	return r
}

func initDBConn() (*sql.DB,error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",os.Getenv("DB_USER"),os.Getenv("DB_PASSWORD"),os.Getenv("DB_NAME"))
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil,err
    }

	return db,nil
}

func main() {

	err := godotenv.Load()
	if err != nil {
        log.Fatal("Error loading .env file")
    }

	db,err := initDBConn()

	// manual dependency injection, can use other library like facebookgo/inject for future development
	movieCore := core.NewMovieCore(db)
	movieServices := services.NewMovieServices(movieCore)
	MovieHandler := handler.NewMovieHandler(movieServices)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	r := setupRouter(MovieHandler)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8081")
}