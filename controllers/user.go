package controllers

import (
//	"io"
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"net/http"
//	"os"
//	"path"
	"github.com/gouthamve/bins/models"
	"github.com/labstack/echo"
)


func CreateUser(c echo.Context) error {
  u := &models.User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	hasher := md5.New()
	hasher.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(hasher.Sum(nil))
	fmt.Println(u)

	err := users.Insert(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}
