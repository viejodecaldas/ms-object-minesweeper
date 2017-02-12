//go:generate goagen bootstrap -d github.com/viejodecaldas/ms-object-minesweeper/src/web/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/viejodecaldas/ms-object-minesweeper/src/web/app"
)

func main() {
	// Create service
	service := goa.New("minesweeper")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "build" controller
	c := NewBuildController(service)
	app.MountBuildController(service, c)
	// Mount "minesweeper" controller
	c2 := NewMinesweeperController(service)
	app.MountMinesweeperController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
