package queries

const CREATE_MOVIE_QUERY = `INSERT INTO movies(title,description,duration,artists,genre,watchURL,vote,viewcount) VALUES($1,$2,$3,$4,$5,$6,$7,$8)`
const UPDATE_MOVIE_QUERY = `UPDATE movies SET title=$1,description=$2,duration=$3,artists=$4,genre=$5,watchURL=$6,vote=$7,viewCount=$8 WHERE id=$9`
const GET_MOVIE_BY_ID_QUERY = `SELECT title,description,duration,artists,genre,watchURL,vote,viewcount FROM movies WHERE id = $1`
const GET_ONE_MOVIE_SORTED_BY_QUERY = "SELECT title,description,duration,artists,genre,watchURL,vote,viewcount from movies ORDER BY %s %s limit 1"
const GET_MOST_VIEWED_GENRE_QUERY = `select genre, sum(viewcount) as totalView from movies group by genre order by totalview desc limit 1`

const GET_ALL_MOVIES_PAGINATION_QUERY = "SELECT * FROM movies ORDER BY id LIMIT %d OFFSET (%d-1) * %d;"
const GET_MOVIES_BY_QUERY = "SELECT * FROM movies where %s"
