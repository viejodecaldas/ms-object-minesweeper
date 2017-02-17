package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("minesweeper", func() {
	BasePath("/minesweeper")
	DefaultMedia(BoardType)

	//Minesweeper group
	Action("StartNewGame", func() {
		Description("Sets the new board.")
		Routing(GET("/new-game"))
		Params(func() {
			Param("width", Integer, "Board width.")
			Param("height", Integer, "Board height.")
			Param("mines", Integer, "Amaount of mines that will be set on the board.")
			Required("width", "height", "mines")
		})
		Response(OK, BoardType)
		Response(NotFound)
	})

	Action("ClickedCell", func() {
		Description("Analyze cell clicked.")
		Routing(PUT("/clicked-cell/:row/:cell"))
		Response(Accepted, BoardType)
		Response(OK, ConfirmationMedia)
		Response(NotFound)
		//Get payload manually from body request
	})

})

// BoardType defines the media type used to render the minesweeper board.
var BoardType = MediaType("application/vnd.goa.boardtype+json", func() {
	Description("Minesweeper's board structure.")
	Attributes(func() {
		Attribute("mineNum", Integer, "Number of mines set.")
		Attribute("width", Integer, "Board width.")
		Attribute("height", Integer, "Board height.")
		Attribute("cellsRemaining", Integer, "Number of cells remaining untill game ends.")
		Attribute("grid", ArrayOf(ArrayOf(CellType)), "Board grid.")
		Required("mineNum", "width", "height", "grid", "cellsRemaining")
	})

	View("default", func() {
		Attribute("mineNum")
		Attribute("width")
		Attribute("height")
		Attribute("cellsRemaining")
		Attribute("grid")
	})
})

// CellType defines the media type used to handle cell behaviour.
var CellType = MediaType("application/vnd.goa.celltype+json", func() {
	Description("Minesweeper's cell structure")
	Attributes(func() {
		Attribute("mine", Boolean, "Attribute to know if cell has a mine or not.")
		Attribute("clicked", Boolean, "Attribute to know if cell has been clicked.")
		Attribute("value", Integer, "Attribute to know value of adjeacent cells.")
		Required("mine", "clicked", "value")
	})

	View("default", func() {
		Attribute("mine")
		Attribute("clicked")
		Attribute("value")
	})
})