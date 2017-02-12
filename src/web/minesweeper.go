package main

import (
	"github.com/goadesign/goa"
	"github.com/viejodecaldas/ms-object-minesweeper/src/web/app"
)

// MinesweeperController implements the minesweeper resource.
type MinesweeperController struct {
	*goa.Controller
}

// NewMinesweeperController creates a minesweeper controller.
func NewMinesweeperController(service *goa.Service) *MinesweeperController {
	return &MinesweeperController{Controller: service.NewController("MinesweeperController")}
}

// StartNewGame runs the StartNewGame action.
func (c *MinesweeperController) StartNewGame(ctx *app.StartNewGameMinesweeperContext) error {
	// MinesweeperController_StartNewGame: start_implement

	// Put your logic here

	// MinesweeperController_StartNewGame: end_implement
	res := &app.GoaBoardtype{}
	return ctx.OK(res)
}
