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
)

func AddtoPlaylist(c echo.Context) error {
	p := &models.Playlist{}
	if err := c.Bind(p); err != nil {
		return err
	}




	err := playlists.Insert(p.Song)
	if err != nil {
		return err
	}


}
