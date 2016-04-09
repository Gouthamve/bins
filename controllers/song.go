package controllers

import (
	"net/http"

	"github.com/gouthamve/bins/models"
	"github.com/labstack/echo"
)

// CreateSong is POST Handler for Song
func CreateSong(c echo.Context) error {
	newS := &models.Song{}
	if err := c.Bind(newS); err != nil {
		return err
	}

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
