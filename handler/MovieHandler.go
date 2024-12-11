package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"test-msbu/models"
	"test-msbu/services"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	MovieServices services.IMovieServices
}

type IMovieHandler interface {
	CreateMovie(c *gin.Context)
	UpdateMovieById(c *gin.Context)
	GetMostViewedMovie(c *gin.Context)
	GetMostViewedGenre(c *gin.Context)
	GetMoviesPaginated(c *gin.Context)
}

func NewMovieHandler(ms services.IMovieServices) IMovieHandler {
	return &MovieHandler{
		MovieServices: ms,
	}
}

func (h *MovieHandler) CreateMovie(c *gin.Context) {
	var movie models.Movie

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.MovieServices.CreateMovie(movie)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": fmt.Sprintf("Failed to create movie because of %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Movie created successfully",
	})
}

func (h *MovieHandler) UpdateMovieById(c *gin.Context) {
	var movie models.Movie
	movieId := c.Param("id")

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.MovieServices.UpdateMovieById(movie, movieId)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": fmt.Sprintf("Failed to update movie because of : %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Movie updated successfully",
	})
}

func (h *MovieHandler) GetMostViewedMovie(c *gin.Context) {

	movie, err := h.MovieServices.GetMostViewedMovie()

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": fmt.Sprintf("Failed to get most viewed movie because of : %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":       movie.Title,
		"description": movie.Description,
		"duration":    movie.Duration,
		"artists":     movie.Artists,
		"genre":       movie.Genre,
		"watchURL":    movie.WatchURL,
		"vote":        movie.Vote,
		"viewCount":   movie.ViewCount,
	})
}

func (h *MovieHandler) GetMostViewedGenre(c *gin.Context) {

	genre, viewCount, err := h.MovieServices.GetMostViewedGenre()

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": fmt.Sprintf("Failed to get most viewed genre because of : %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"genre":     genre,
		"viewCount": viewCount,
	})
}

func (mh *MovieHandler) GetMoviesPaginated(c *gin.Context) {

	page := c.Params.ByName("page")
	perPage := c.Params.ByName("perPage")

	pageNum, err := strconv.Atoi(page)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Failed to get movies because of %s", err.Error()),
		})
		return
	}

	perPageNum, err := strconv.Atoi(perPage)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Failed to get movies because of %s", err.Error()),
		})
		return
	}

	if pageNum < 0 || perPageNum < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "page and perPage params must be greater than 0",
		})
		return
	}

	movies, err := mh.MovieServices.GetMoviesPaginated(pageNum, perPageNum)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("failed to get paginated page because of %s", err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, movies)
}
