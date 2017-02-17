package models

import (
	"errors"
	"math/rand"
)

//Struct to handle the board behaviour
type Board struct {
	Grid [][]Cell
	Height int
	Width int
	MineNum int
	CellsRemaining int
}

//Method to build up the board with parameters sent
func (b *Board) BuildBoard() (err error) {
	//Initialize board and validate values sent
	height := b.Height
	width := b.Width
	mineNum := b.MineNum
	b.CellsRemaining = (b.Height * b.Width) - b.MineNum
	b.Grid = make([][]Cell, height)
	cells := make([]Cell, (height * width))

	if mineNum > b.Height * b.Width - 1 {
		err = errors.New("Number of mines cannot exceed number of cells.")
		return err
	}
	b.MineNum = mineNum
	mines := b.genMines()

	//loop through all cells in the slice, keeping track of index
	for index := range cells {
		cells[index] = Cell{
			Mine: mines[index],
			Clicked: false,
			Value: 0,
		}
	}

	//slice the cell slice up into rows
	for row := range b.Grid {
		b.Grid[row] = cells[(width * row):(width * (row + 1))]
	}

	//go through each element in each row and check if cell has a mine
	for i, row := range b.Grid {
		for j, cell := range row {
			if cell.Mine {
				if b.CheckPosition(i-1, j-1){
					b.Grid[i-1][j-1].Value++
				}
				if b.CheckPosition(i-1, j){
					b.Grid[i-1][j].Value++
				}
				if b.CheckPosition(i-1, j+1){
					b.Grid[i-1][j+1].Value++
				}
				if b.CheckPosition(i, j-1) {
					b.Grid[i][j-1].Value++
				}
				if b.CheckPosition(i, j+1) {
					b.Grid[i][j+1].Value++
				}
				if b.CheckPosition(i+1, j-1) {
					b.Grid[i+1][j-1].Value++
				}
				if b.CheckPosition(i+1, j) {
					b.Grid[i+1][j].Value++
				}
				if b.CheckPosition(i+1, j+1) {
					b.Grid[i+1][j+1].Value++
				}
			}
		}
	}
	return err
}

//Method to get a random position for mines
func (b *Board) genMines() []bool {
	mineIndexList := make([]bool, (b.Height * b.Width))
	i := 0
	for i < b.MineNum {
		index := rand.Intn(b.Height * b.Width)
		if mineIndexList[index] == false {
			mineIndexList[index] = true
			i++
		}
	}
	return mineIndexList
}

//Method to check if position is valid
func (b *Board) CheckPosition(r, c int) bool{
	if r < 0 || r > b.Height-1 {
		return false
	}
	if c < 0 || c > b.Width-1 {
		return false
	}
	return true
}