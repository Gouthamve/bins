package controllers

import (
	"net/http"
	"fmt"
//	"path"
	"github.com/gouthamve/bins/models"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func AddtoPlaylist(c echo.Context) error {
	p := &models.Playlist{}
	if err := c.Bind(p); err != nil {
		return err
	}
	fmt.Println(p.Songs)
	p2 := models.Playlist{}
	err := playlists.Find(bson.M{"name": p.Name}).One(&p2)
	if err != nil {
		//return err
	}

	if p2.Name != "" {
	p2.Songs = append(p2.Songs, p.Songs[0])
	err = playlists.Update(bson.M{"name": p.Name}, bson.M{"$set": bson.M{"Songs": p2.Songs}})
	if err != nil {
		return err
	}
	} else {
		p2 = *p
		err = playlists.Insert(p)
		if err != nil {
			return err
		}
	}


	return c.JSON(http.StatusOK, p2)
}
