package core

import (
	"database/sql"
	"fmt"
	"test-msbu/models"
	"test-msbu/queries"
)

type IMovieCore interface {
	CreateMovie(models.Movie) error
	UpdateMovieById(movie models.Movie, id string) error
	GetMovieById(id string) (movie models.Movie,err error)
	GetOneMovieSortedBy(sortedByParams []string, sortDir string) (movie models.Movie,err error)
}

type MovieCore struct {
    db *sql.DB
}

func NewMovieCore(db *sql.DB) IMovieCore{
	return &MovieCore{
		db,
	}
}

func (mc *MovieCore) CreateMovie(movie models.Movie) error {

	// we put votes & view count to always 0 since it is a new movie
	_,err := mc.db.Exec(queries.CREATE_MOVIE_QUERY,movie.Title,movie.Description,movie.Duration,movie.Artists,movie.Genres,movie.WatchURL,0,0)

	if err != nil {
		return fmt.Errorf("error when executing create movie query to db because of %s",err.Error())
	}
	
	return nil
}

func (mc *MovieCore) UpdateMovieById(movie models.Movie, id string) error {

	_,err := mc.db.Exec(queries.UPDATE_MOVIE_QUERY,movie.Title,movie.Description,movie.Duration,movie.Artists,movie.Genres,movie.WatchURL,movie.Vote,movie.ViewCount, id)

	if err != nil {
		return fmt.Errorf("error when executing update movie query to db because of %s",err.Error())
	}

	return nil
}

func (mc *MovieCore) GetMovieById(id string) (models.Movie,error){

	var title,description,duration,artists,genres,watchURL string
	var vote,viewcount int64

    err := mc.db.QueryRow(queries.GET_MOVIE_BY_ID_QUERY, id).Scan(&title, &description, &duration, &artists, &genres, &watchURL, &vote, &viewcount)

    if err != nil {
       return models.Movie{},err
    }

	return models.Movie{
		Title: title,
		Description: description,
		Duration: duration,
		Artists:artists,
		Genres: genres,
		WatchURL: watchURL,
		Vote: vote,
		ViewCount: viewcount,
	},nil
}

func (mc *MovieCore) GetOneMovieSortedBy(sortedBy []string, sortDir string) (models.Movie,error) {

	var title,description,duration,artists,genres,watchURL string
	var vote,viewcount int64

	var unifiedSortedByParams string 
	for _,v := range sortedBy {
		unifiedSortedByParams += v +","
	}

	unifiedSortedByParams = unifiedSortedByParams[:len(unifiedSortedByParams)-1]
	query := fmt.Sprintf(queries.GET_ONE_MOVIE_SORTED_BY_QUERY,unifiedSortedByParams,sortDir)
    err := mc.db.QueryRow(query).Scan(&title, &description, &duration, &artists, &genres, &watchURL, &vote, &viewcount)

    if err != nil {
       return models.Movie{},err
    }

	return models.Movie{
		Title: title,
		Description: description,
		Duration: duration,
		Artists:artists,
		Genres: genres,
		WatchURL: watchURL,
		Vote: vote,
		ViewCount: viewcount,
	},nil

}