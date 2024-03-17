package models

type Movie struct {
	MovieID     int     `json:"movieId"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ReleaseDate string  `json:"releaseDate"`
	Reting      float32 `json:"rating"`
	Actors      []Actor `json:"actors"`
}
