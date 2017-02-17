package convert

import (
	"github.com/viejodecaldas/ms-object-minesweeper/src/web/app"
	"github.com/viejodecaldas/ms-object-minesweeper/src/app/models"
)

//Function to parse structure from model to response for the client
func ToGoaBoard(board models.Board) *app.GoaBoardtype {
	//Sets first cell inside the board
	boardCells := make([][]*app.GoaCelltype, board.Height)

	//Iterate through every cell in the board an map it into the
	//proper response.
	for rowIndex, rowValue := range board.Grid {

		cells := make([]*app.GoaCelltype, 0, 0)
		for _, cellValue := range rowValue {
			cell := &app.GoaCelltype{
				Value: cellValue.Value,
				Mine: cellValue.Mine,
				Clicked: cellValue.Clicked,
			}

			//Set every value into the cell array for the current row
			cells = append(cells, cell)
		}

		boardCells[rowIndex] = cells
	}

	//Set values to the response structure
	goa := &app.GoaBoardtype{
		Height: board.Height,
		Width: board.Width,
		MineNum: board.MineNum,
		CellsRemaining: board.CellsRemaining,
		Grid: boardCells,
	}

	return goa
}

//Function to parse from Goa structure to model structure
func FromGoaBoard(board *app.GoaBoardtype) models.Board {
	//Sets first cell inside the board
	boardCells := make([][]models.Cell, board.Height)

	//Iterate through every cell in the board an map it into the
	//proper response.
	for rowIndex, rowValue := range board.Grid {

		cells := make([]models.Cell, 0, 0)
		for _, cellValue := range rowValue {
			cell := models.Cell{
				Value: cellValue.Value,
				Mine: cellValue.Mine,
				Clicked: cellValue.Clicked,
			}

			//Set every value into the cell array for the current row
			cells = append(cells, cell)
		}

		boardCells[rowIndex] = cells
	}

	//Sets the values for the models structure
	model := models.Board{
		Height: board.Height,
		Width: board.Width,
		MineNum: board.MineNum,
		CellsRemaining: board.CellsRemaining,
		Grid: boardCells,
	}

	return model
}