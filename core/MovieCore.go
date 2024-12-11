package core

import (
	"database/sql"
	"fmt"
	"strconv"
	"test-msbu/models"
	"test-msbu/queries"
)

type IMovieCore interface {
	CreateMovie(models.Movie) error
	UpdateMovieById(movie models.Movie, id string) error
	GetMovieById(id string) (movie models.Movie, err error)
	GetOneMovieSortedBy(sortedByParams []string, sortDir string) (movie models.Movie, err error)
	GetMostViewedGenre() (genre string, viewCount int, err error)
	GetMoviesPaginated(page, perPage int) (movies []models.Movie, err error)
	GetMoviesByOptions(map[string]string) (movies []models.Movie, err error)
}

type MovieCore struct {
	db *sql.DB
}

func NewMovieCore(db *sql.DB) IMovieCore {
	return &MovieCore{
		db,
	}
}

func (mc *MovieCore) CreateMovie(movie models.Movie) error {

	// we put votes & view count to always 0 since it is a new movie
	_, err := mc.db.Exec(queries.CREATE_MOVIE_QUERY, movie.Title, movie.Description, movie.Duration, movie.Artists, movie.Genre, movie.WatchURL, 0, 0)

	if err != nil {
		return fmt.Errorf("error when executing create movie query to db because of %s", err.Error())
	}

	return nil
}

func (mc *MovieCore) UpdateMovieById(movie models.Movie, id string) error {

	_, err := mc.db.Exec(queries.UPDATE_MOVIE_QUERY, movie.Title, movie.Description, movie.Duration, movie.Artists, movie.Genre, movie.WatchURL, movie.Vote, movie.ViewCount, id)

	if err != nil {
		return fmt.Errorf("error when executing update movie query to db because of %s", err.Error())
	}

	return nil
}

func (mc *MovieCore) GetMovieById(id string) (models.Movie, error) {

	var title, description, duration, artists, genre, watchURL string
	var vote, viewcount int64

	err := mc.db.QueryRow(queries.GET_MOVIE_BY_ID_QUERY, id).Scan(&title, &description, &duration, &artists, &genre, &watchURL, &vote, &viewcount)

	if err != nil {
		return models.Movie{}, err
	}

	return models.Movie{
		Title:       title,
		Description: description,
		Duration:    duration,
		Artists:     artists,
		Genre:       genre,
		WatchURL:    watchURL,
		Vote:        vote,
		ViewCount:   viewcount,
	}, nil
}

func (mc *MovieCore) GetOneMovieSortedBy(sortedBy []string, sortDir string) (models.Movie, error) {

	var title, description, duration, artists, genre, watchURL string
	var vote, viewcount int64

	var unifiedSortedByParams string
	for _, v := range sortedBy {
		unifiedSortedByParams += v + ","
	}

	unifiedSortedByParams = unifiedSortedByParams[:len(unifiedSortedByParams)-1]
	query := fmt.Sprintf(queries.GET_ONE_MOVIE_SORTED_BY_QUERY, unifiedSortedByParams, sortDir)
	err := mc.db.QueryRow(query).Scan(&title, &description, &duration, &artists, &genre, &watchURL, &vote, &viewcount)

	if err != nil {
		return models.Movie{}, err
	}

	return models.Movie{
		Title:       title,
		Description: description,
		Duration:    duration,
		Artists:     artists,
		Genre:       genre,
		WatchURL:    watchURL,
		Vote:        vote,
		ViewCount:   viewcount,
	}, nil
}

func (mc *MovieCore) GetMostViewedGenre() (string, int, error) {

	var genre string
	var viewCount int
	err := mc.db.QueryRow(queries.GET_MOST_VIEWED_GENRE_QUERY).Scan(&genre, &viewCount)

	if err != nil {
		return "", 0, err
	}

	return genre, viewCount, nil
}

func (mc *MovieCore) GetMoviesPaginated(page, perPage int) ([]models.Movie, error) {

	query := fmt.Sprintf(queries.GET_ALL_MOVIES_PAGINATION_QUERY, perPage, page, perPage)
	rows, err := mc.db.Query(query)

	var movies []models.Movie
	var title, description, duration, artists, genre, watchURL string
	var id, vote, viewcount int64

	for rows.Next() {
		err = rows.Scan(&id, &title, &description, &duration, &artists, &genre, &watchURL, &vote, &viewcount)

		if err != nil {
			return []models.Movie{}, err
		}

		movies = append(movies, models.Movie{
			Id:          id,
			Title:       title,
			Description: description,
			Duration:    duration,
			Artists:     artists,
			Genre:       genre,
			WatchURL:    watchURL,
			Vote:        vote,
			ViewCount:   viewcount,
		})
	}

	return movies, err
}

func (mc *MovieCore) GetMoviesByOptions(searchOpts map[string]string) ([]models.Movie, error) {

	var whereCond, query string
	var searchValue []interface{}

	if len(searchOpts) == 0 {
		query = fmt.Sprintf(queries.GET_MOVIES_BY_QUERY, "1=1")
	} else {
		counter := 1
		for k, v := range searchOpts {
			whereCond += k + " LIKE $" + strconv.Itoa(counter) + " AND "
			searchValue = append(searchValue, "%"+v+"%")
			counter++
		}

		whereCond = whereCond[:len(whereCond)-4]
		fmt.Println(searchOpts)

		query = fmt.Sprintf(queries.GET_MOVIES_BY_QUERY, whereCond)
	}

	rows, err := mc.db.Query(query, searchValue...)

	if err != nil {
		return []models.Movie{}, err
	}

	var movies []models.Movie
	var title, description, duration, artists, genre, watchURL string
	var id, vote, viewcount int64

	for rows.Next() {

		err = rows.Scan(&id, &title, &description, &duration, &artists, &genre, &watchURL, &vote, &viewcount)
		if err != nil {
			return []models.Movie{}, err
		}

		movies = append(movies, models.Movie{
			Id:          id,
			Title:       title,
			Description: description,
			Duration:    duration,
			Artists:     artists,
			Genre:       genre,
			WatchURL:    watchURL,
			Vote:        vote,
			ViewCount:   viewcount,
		})
	}

	return movies, err
}
