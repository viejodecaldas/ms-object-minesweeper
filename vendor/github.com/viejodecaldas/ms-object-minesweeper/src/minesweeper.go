package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/viejodecaldas/ms-object-minesweeper/src/app/app"
	"github.com/viejodecaldas/ms-object-minesweeper/src/app/models"
	"strconv"
)

// TicketController returns information about a ticket.
type MinesweeperController struct{}

// Mount setups the routes for the EventsController.
func (mc MinesweeperController) Mount(e *echo.Echo) {
	g := e.Group("/minesweeper")

	g.GET("/new-game", mc.StartNewGame)
	g.PUT("/clicked-cell/:row/:cell", mc.ClickedCell)
	g.POST("/save-game", mc.SaveGame)
	g.GET("/load-game/:gameID", mc.LoadGame)
}

func (ms *MinesweeperController) LoadGame(ctx echo.Context) error {
	//Validate parameters sent
	gameID, err := strconv.Atoi(ctx.Param("gameID"))
	if err != nil {
		return app.Error(ctx,
			fmt.Errorf("Could not parse game ID value %s. Error: %s",
				ctx.QueryParam("gameID"),
				err.Error()))
	}

	var board models.Board
	board.LoadGame(gameID)

	//Return the saved board game
	return app.Success(ctx, board)
}

//Method to save a board game
func (ms *MinesweeperController) SaveGame(ctx echo.Context) error {
	//Decode request body
	var board models.Board
	err := ctx.Bind(&board)
	if err != nil {
		return app.Error(ctx,
			fmt.Errorf("Could not decode body from save game endpoint. Error: %s",
				err.Error()))
	}

	//Save the actual state of the board
	err = board.SaveGame()
	if err != nil {
		return app.Error(ctx,
			fmt.Errorf("Could not save game. Error: %s",
				err.Error()))
	}

	return app.OK(ctx, "Game Saved Successfully!")
}

// ClickedCell runs the ClickedCell action.
func (ms *MinesweeperController) ClickedCell(ctx echo.Context) error {
	// MinesweeperController_ClickedCell: start_implement

	//Validate parameters sent
	row, err := strconv.Atoi(ctx.Param("row"))
	if err != nil {
		return app.Error(ctx,
			fmt.Errorf("Could not parse row value %s. Error: %s",
				ctx.QueryParam("row"),
				err.Error()))
	}

	cell, err := strconv.Atoi(ctx.Param("cell"))
	if err != nil {
		return app.Error(ctx,
			fmt.Errorf("Could not parse cell value %s. Error: %s",
				ctx.QueryParam("cell"),
				err.Error()))
	}

	//Decode request body
	var board models.Board
	err = ctx.Bind(&board)
	if err != nil {
		return app.Error(ctx,
			fmt.Errorf("Could not decode body from clicked cell endpoint. Error: %s",
				err.Error()))
	}

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
			board.CellsRemaining--
		case models.AlreadyClicked:
			fmt.Println("you've already clicked that position")
		}
	}

	//If user clicked on a cell that had a mine then the game is over.
	if !notLost {
		return app.Error(ctx,
			fmt.Errorf("Sorry, you lost!"))
	}

	//If user hasn't find a mine and there is no more cells remaining
	//Then user has won the game
	if notLost && board.CellsRemaining == 0 {
		return app.OK(ctx, "Congratulations, you won!")
	}

	//Return the board with the new set of values
	return app.Success(ctx, board)
	// MinesweeperController_ClickedCell: end_implement
}

// StartNewGame runs the StartNewGame action.
func (ms *MinesweeperController) StartNewGame(ctx echo.Context) error {
	// MinesweeperController_StartNewGame: start_implement

	width, err := strconv.Atoi(ctx.QueryParam("width"))
	if err != nil {
		return app.Error(ctx,
			fmt.Errorf("could not parse width param %s. Error: %s",
				ctx.QueryParam("width"),
				err.Error()))
	}

	height, err := strconv.Atoi(ctx.QueryParam("height"))
	if err != nil {
		return app.Error(ctx,
			fmt.Errorf("could not parse height param %s. Error: %s",
				ctx.QueryParam("height"),
				err.Error()))
	}

	mines, err := strconv.Atoi(ctx.QueryParam("mines"))
	if err != nil {
		return app.Error(ctx,
			fmt.Errorf("could not parse mines param %s. Error: %s",
				ctx.QueryParam("mines"),
				err.Error()))
	}

	//Set the initial values on the board
	var board = models.Board{
		Width:   width,
		Height:  height,
		MineNum: mines,
	}

	//Build the board with initial setup
	err = board.BuildBoard()
	if err != nil {
		return app.Error(ctx,
			fmt.Errorf("Error: %s",
				err.Error()))
	}

	//Return the initial board
	return app.Success(ctx, board)

	// MinesweeperController_StartNewGame: end_implement
}
