package convert

import (
	"github.com/viejodecaldas/ms-object-minesweeper/src/app/models"
	"github.com/viejodecaldas/ms-object-minesweeper/src/web/app"
)

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

	goa := &app.GoaBoardtype{
		Height: board.Height,
		Width: board.Width,
		MineNum: board.MineNum,
		Grid: boardCells,
	}

	return goa
}

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

	model := models.Board{
		Height: board.Height,
		Width: board.Width,
		MineNum: board.MineNum,
		Grid: boardCells,
	}

	return model
}