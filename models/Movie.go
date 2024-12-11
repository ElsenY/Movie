package models

type Movie struct {
	Id          int
	Title       string
	Description string
	Duration    string
	Artists     string
	Genre       string
	WatchURL    string
	Vote        int
	ViewCount   int
}
