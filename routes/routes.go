package routes

import (
	"net/http"
	"os"
	"test-msbu/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(mh handler.IMovieHandler) *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("movies", mh.GetMoviesPaginated)
	r.GET("movieswithopts", mh.GetMoviesByOptions)
	r.GET("movieviewcount", mh.GetMovieViewCount)

	// basic auth for admin route
	admin := r.Group("/", gin.BasicAuth(gin.Accounts{
		os.Getenv("ADMIN_USERNAME"): os.Getenv("ADMIN_PASSWORD"),
	}))

	admin.POST("movie", mh.CreateMovie)
	admin.PUT("movie/:id", mh.UpdateMovieById)
	admin.GET("mostviewedmovie", mh.GetMostViewedMovie)
	admin.GET("mostviewedgenre", mh.GetMostViewedGenre)

	return r
}
