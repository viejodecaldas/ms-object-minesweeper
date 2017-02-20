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
	Mine    bool `json:"mine"`
	Clicked bool `json:"clicked"`
	Value   int  `json:"value"`
}

//Method to know how to render the cell depending on the attributes values
func (c *Cell) Render() string {
	switch {
	case c.Clicked == false:
		return "?"
	case c.Mine == true:
		return "X"
	default:
		return strconv.Itoa(c.Value)
	}
}

//Method that return the cell's behaviour
func (c *Cell) Click() int {
	switch {
	case c.Clicked == true:
		return AlreadyClicked
	case c.Mine == false:
		return Ok
	case c.Mine == true:
		return Mine
	}
	return AlreadyClicked
}
