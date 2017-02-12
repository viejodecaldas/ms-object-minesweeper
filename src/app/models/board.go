package models

type Board struct {
	grid [][]*Cell
	height int
	width int
	mineNum int
}

