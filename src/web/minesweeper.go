package main

import (
	"github.com/goadesign/goa"
	"github.com/viejodecaldas/ms-object-minesweeper/src/web/app"
	"fmt"
	"encoding/json"
	"strconv"
	"github.com/viejodecaldas/ms-object-minesweeper/src/app/convert"
	"github.com/viejodecaldas/ms-object-minesweeper/src/app/models"
)

// MinesweeperController implements the minesweeper resource.
type MinesweeperController struct {
	*goa.Controller
}

// NewMinesweeperController creates a minesweeper controller.
func NewMinesweeperController(service *goa.Service) *MinesweeperController {
	return &MinesweeperController{Controller: service.NewController("MinesweeperController")}
}

// ClickedCell runs the ClickedCell action.
func (c *MinesweeperController) ClickedCell(ctx *app.ClickedCellMinesweeperContext) error {
	// MinesweeperController_ClickedCell: start_implement

	//Validate parameters sent
	row, err := strconv.Atoi(ctx.Row)
	if err != nil {
		fmt.Println("Could not parse row value.")
		return goa.ErrInternal(err)
	}

	cell, err := strconv.Atoi(ctx.Cell)
	if err != nil {
		fmt.Println("Could not parse cell value.")
		return goa.ErrInternal(err)
	}

	//Decode request body
	var goaBoard app.GoaBoardtype
	err = json.NewDecoder(ctx.Body).Decode(&goaBoard)
	if err != nil {
		fmt.Println("Could not decode body from clicked cell endpoint.")
		return goa.ErrInternal(err)
	}

	//Parse object to model to handle the game
	board := convert.FromGoaBoard(&goaBoard)

	//Check position clicked by the user and validate
	var notLost = true
	if board.CheckPosition(row, cell) {
		clickedCell := board.Grid[row][cell]
		switch clickedCell.Click() {
		case models.Mine:
			clickedCell.Clicked = true
			notLost = false
		case models.Ok:
			clickedCell.Clicked = true
			board.Grid[row][cell] = clickedCell
		case models.AlreadyClicked:
			fmt.Println("you've already clicked that position")
		}
	}

	//If user clicked on a cell that had a mine then the game is over.
	if !notLost {
		return goa.ErrInternal(fmt.Errorf("Sorry, you lost!"))
	}

	//Return the board with the new set of values
	return ctx.OK(convert.ToGoaBoard(board))
	// MinesweeperController_ClickedCell: end_implement
}

// StartNewGame runs the StartNewGame action.
func (c *MinesweeperController) StartNewGame(ctx *app.StartNewGameMinesweeperContext) error {
	// MinesweeperController_StartNewGame: start_implement

	//Set the initial values on the board
	var board = models.Board{
		Width: ctx.Width,
		Height: ctx.Height,
		MineNum: ctx.Mines,
	}

	//Build the board with initial setup
	err := board.BuildBoard()
	if err != nil {
		fmt.Println("Error", err.Error())
		return goa.ErrInternal(err)
	}

	//Return the initial board
	return ctx.OK(convert.ToGoaBoard(board))

	// MinesweeperController_StartNewGame: end_implement
}
