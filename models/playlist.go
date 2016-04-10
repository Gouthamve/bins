package models

// Song contains all the feilds of a song
type Playlist struct {
	Name      string `json:"name"`
	Songs     []Song `json:"songs"`
}
