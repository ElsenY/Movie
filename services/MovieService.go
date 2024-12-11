package services

import (
	"database/sql"
	"errors"
	"fmt"
	"test-msbu/core"
	"test-msbu/models"
)

type IMovieServices interface {
	CreateMovie(models.Movie) error
	UpdateMovieById(movie models.Movie, id string) error
	GetMostViewedMovie() (movie models.Movie, err error)
	GetMostViewedGenre() (genre string, viewCount int, err error)
	GetMoviesPaginated(page, perPage int) (movies []models.Movie, err error)
	GetMoviesByOptions(searchOpts map[string]string) (movies []models.Movie, err error)
	GetMovieViewCount(id string) (title string, viewCount int, err error)
}

type MovieServices struct {
	MovieCore core.IMovieCore
}

func NewMovieServices(mc core.IMovieCore) IMovieServices {

	return &MovieServices{
		MovieCore: mc,
	}
}

func (ms *MovieServices) CreateMovie(movie models.Movie) error {

	err := ms.MovieCore.CreateMovie(movie)

	if err != nil {
		return err
	}

	return nil
}

func (ms *MovieServices) UpdateMovieById(movie models.Movie, id string) error {

	// check if movie exist
	_, err := ms.MovieCore.GetMovieById(id)

	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("no movie found for specified id")
		} else {
			return fmt.Errorf("there is an error while searching for the movie by id becase of %s", err)
		}
	}

	// if exist, we update the movie
	err = ms.MovieCore.UpdateMovieById(movie, id)

	if err != nil {
		return err
	}

	return nil
}

func (ms *MovieServices) GetMostViewedMovie() (models.Movie, error) {

	movie, err := ms.MovieCore.GetOneMovieSortedBy([]string{"viewcount"}, "DESC")

	if err != nil {
		return models.Movie{}, fmt.Errorf("there is an error while searching for most viewed movie becase of %s", err)
	}

	return movie, nil
}

func (ms *MovieServices) GetMostViewedGenre() (string, int, error) {

	genre, viewCount, err := ms.MovieCore.GetMostViewedGenre()

	if err != nil {
		return "", 0, err
	}

	return genre, viewCount, nil
}

func (ms *MovieServices) GetMoviesPaginated(page, perPage int) ([]models.Movie, error) {

	movies, err := ms.MovieCore.GetMoviesPaginated(page, perPage)

	if err != nil {
		return []models.Movie{}, err
	}

	return movies, nil
}

func (ms *MovieServices) GetMoviesByOptions(searchOpts map[string]string) ([]models.Movie, error) {
	movies, err := ms.MovieCore.GetMoviesByOptions(searchOpts)

	if err != nil {
		return []models.Movie{}, err
	}

	return movies, nil
}

func (ms *MovieServices) GetMovieViewCount(id string) (string, int, error) {

	movie, err := ms.MovieCore.GetMovieById(id)

	if err != nil {
		return "", 0, err
	}

	return movie.Title, movie.ViewCount, nil
}
