package models

import "strconv"

//Const to control possible cell state
const (
	AlreadyClicked int = iota //Const to notify the client the cell was already clicked
	Ok
	Mine
)

//Defines the cell behaviour
type Cell struct {
	mine bool
	clicked bool
	value int
}

//Method to know how to render the cell depending on the attributes values
func (c *Cell) Render() string {
	switch {
	case c.clicked == false:
		return "?"
	case c.mine == true:
		return "X"
	default:
		return strconv.Itoa(c.value)
	}
}

//Method that return the cell's behaviour
func (c *Cell) Click() int {
	switch {
	case c.clicked == true:
		return AlreadyClicked
	case c.mine == false:
		return Ok
	case c.mine == true:
		return Mine
	}
	return AlreadyClicked
}