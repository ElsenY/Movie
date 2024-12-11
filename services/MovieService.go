package services

import (
	"database/sql"
	"errors"
	"fmt"
	"test-msbu/core"
	"test-msbu/models"
)

type IMovieServices interface{
	CreateMovie(models.Movie) error
	UpdateMovieById(movie models.Movie, id string) error
	GetMostViewedMovie() (movie models.Movie,err error)
	GetMostViewedGenre() (genre string, viewCount int, err error)
}

type MovieServices struct {
    MovieCore core.IMovieCore
}

func NewMovieServices(mc core.IMovieCore) IMovieServices{

	return &MovieServices{
		MovieCore: mc,
	}
}

func (ms *MovieServices) CreateMovie(movie models.Movie) error{

	err := ms.MovieCore.CreateMovie(movie)

	if err != nil {
		return err
	}

	return nil
}

func (ms *MovieServices) UpdateMovieById(movie models.Movie, id string) error {

	// check if movie exist
	_,err := ms.MovieCore.GetMovieById(id)

	if err != nil {
        if err == sql.ErrNoRows {
            return errors.New("no movie found for specified id")
        } else {
            return fmt.Errorf("there is an error while searching for the movie by id becase of %s",err)
        }
    }

	// if exist, we update the movie
	err = ms.MovieCore.UpdateMovieById(movie,id)

	if err != nil {
		return err
	}

	return nil
}

func (ms *MovieServices) GetMostViewedMovie() (models.Movie,error) {

	movie,err := ms.MovieCore.GetOneMovieSortedBy([]string{"viewcount"},"DESC")

	if err != nil {
		return models.Movie{},fmt.Errorf("there is an error while searching for most viewed movie becase of %s",err)
	}

	return movie,nil
} 

func (ms *MovieServices) GetMostViewedGenre() (string,int,error){

	genre,viewCount,err := ms.MovieCore.GetMostViewedGenre()

	if err != nil {
		return "",0,err
	}

	return genre, viewCount, nil
}