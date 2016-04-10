package main

import (
	"net/http"


	"github.com/gouthamve/bins/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.SetDebug(true)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.Get("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!\n")
	})

	e.Get("/api/songs", controllers.GetSongs)
	e.Get("/api/:artist/:album/song", controllers.GetSongs)
	e.Post("/api/songs", controllers.CreateSong)
	e.Post("/api/user/add", controllers.CreateUser)
	// Start server
	e.Run(standard.New(":1323"))
}
