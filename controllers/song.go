package controllers

import (
	"io"
	"net/http"
	"os"
	"strings"
//	"path"
	"fmt"
	"github.com/gouthamve/bins/models"
	"github.com/labstack/echo"
	"github.com/mikkyang/id3-go"
)



// CreateSong is POST Handler for Song
func CreateSong(c echo.Context) error {
	newS := &models.Song{}


		form, err := c.MultipartForm()
		if err != nil {
			return err
		}
		files := form.File["songs"]

		for _, file := range files {
			// Source
			src, err := file.Open()
			if err != nil {
				return err
			}
			defer src.Close()



			// Destination
			dst, err := os.Create("test/"+file.Filename)
			if err != nil {
				return err
			}

			// Copy
			if _, err = io.Copy(dst, src); err != nil {
				return err
			}

							f, err := id3.Open("test/"+file.Filename)
							newS.Title = strings.Trim(f.Title(), "\x00")
							newS.Artist = strings.Trim(f.Artist(), "\x00")
							newS.Album = strings.Trim(f.Album(), "\x00")
							newS.Year = f.Year()
							newS.Genre = f.Genre()
							newS.Size = f.Size()
							newS.ID3Version = f.Version()
							dirpath := "Music/"+newS.Artist + "/" + newS.Album
							err = os.MkdirAll(dirpath, 0777)
							if err != nil {
								return err
							}


							 dst.Close()
							f.Close()
							err =  os.Rename("test/"+file.Filename, dirpath+"/"+file.Filename)

          		if err != nil {
								 fmt.Println(err)
								 return err
							 }

							 newS.Location = dirpath+"/"+file.Filename

							insertSong(*newS)



							fmt.Println(newS)


		}


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
