package repository

import "test-msbu/models"

type Movie interface {
	CreateMovie(models.Movie) error
	UpdateMovie(models.Movie) error
	GetMovie(models.Movie) error
	VoteMovie(models.Movie) error
}