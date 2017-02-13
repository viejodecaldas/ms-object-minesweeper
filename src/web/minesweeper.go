package main

import (
	"github.com/goadesign/goa"
	"github.com/viejodecaldas/ms-object-minesweeper/src/web/app"
	"github.com/viejodecaldas/ms-object-minesweeper/src/app/models"
	"fmt"
	"github.com/viejodecaldas/ms-object-minesweeper/src/app/convert"
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

	var board = models.Board{
		Width: ctx.Width,
		Height: ctx.Height,
		MineNum: ctx.Mines,
	}

	err := board.BuildBoard()
	if err != nil {
		fmt.Println("Error", err.Error())
		return goa.ErrInternal(err)
	}

	return ctx.OK(convert.ToGoaBoard(board))
	// MinesweeperController_StartNewGame: end_implement
}
