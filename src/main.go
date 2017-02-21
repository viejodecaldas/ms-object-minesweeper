package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"os"
	"fmt"
)

var (
	BUILD string
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		panic(fmt.Errorf("PORT var is not set."))
	}

	// setup echo instance
	server := echo.New()

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Pre(middleware.RemoveTrailingSlash())

	// Set up build route
	server.GET("/build", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]string{
			"build": BUILD,
		})
	})

	//Mount controller
	minesweeperController := MinesweeperController{}
	minesweeperController.Mount(server)

	// graceful shutdown - setting up the termination timeout to 30 seconds.
	server.Start(":" + port)
}