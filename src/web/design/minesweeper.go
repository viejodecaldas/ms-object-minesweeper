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

})

// BoardType defines the media type used to render the minesweeper board.
var BoardType = MediaType("application/vnd.goa.boardtype+json", func() {
	Description("Minesweeper's board structure.")
	Attributes(func() {
		Attribute("mineNum", Integer, "Number of mines set.")
		Attribute("width", Integer, "Board width.")
		Attribute("height", Integer, "Board height.")
		Attribute("grid", ArrayOf(ArrayOf(CellType)), "Board grid.")
		Required("mineNum", "width", "height")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("mineNum")   // Media types may have multiple views and must
		Attribute("width") // have a "default" view.
		Attribute("height")
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