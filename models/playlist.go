package models

// Song contains all the feilds of a song
type Playlist struct {
	Name      string `json:"name"`
	Songs     []models.Song `json:"songs"`
}
