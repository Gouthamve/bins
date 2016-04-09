package models

// Song contains all the feilds of a song
type Song struct {
	Title      string `json:"title"`
	Artist     string `json:"artist"`
	ArtistID   string `json:"artistId"`
	Album      string `json:"album"`
	AlbumID    string `json:"albumId"`
	Year       string `json:"year"`
	Genre      string `json:"genre"`
	GenreID    string `json:"genreId"`
	Size       int    `json:"size"`
	ID3Version string `json:"id3v"`
}
