package models

type Movie struct {
	Id          int64
	Title       string
	Description string
	Duration    string
	Artists     string
	Genre       string
	WatchURL    string
	Vote        int64
	ViewCount   int64
}
