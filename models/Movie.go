package models

type Movie struct {
	Id int
	Title string
	Description string
	Duration string
	Artists string
	Genres string
	WatchURL string
	Vote int64
	ViewCount int64
}