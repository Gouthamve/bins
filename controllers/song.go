package controllers

import (
	"io"
	"net/http"
	"os"
	"path"

	"github.com/gouthamve/bins/models"
	"github.com/labstack/echo"
	"github.com/mikkyang/id3-go"
)

// CreateSong is POST Handler for Song
func CreateSong(c echo.Context) error {
	newS := &models.Song{}

	file, err := c.FormFile("song")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(path.Join(os.TempDir(), file.Filename))
	if err != nil {
		return err
	}
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	defer dst.Close()

	f, err := id3.Open(dst.Name())
	newS.Title = f.Title()
	newS.Artist = f.Artist()
	newS.Album = f.Album()
	newS.Year = f.Year()
	newS.Genre = f.Genre()
	newS.Size = f.Size()
	newS.ID3Version = f.Version()
	insertSong(*newS)

	return c.JSON(http.StatusCreated, newS)
}

// GetSongs is the GET handler for songs
func GetSongs(c echo.Context) error {
	s := getSongs()
	return c.JSON(http.StatusOK, s)
}

func insertSong(s models.Song) error {
	err := songs.Insert(s)
	if err != nil {
		return err
	}

	return nil
}

func getSongs() (s []models.Song) {
	songs.Find(nil).All(&s)
	return
}

func deleteSong(id string) (err error) {
	err = songs.RemoveId(id)
	return
}
