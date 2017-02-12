package design

import (
. "github.com/goadesign/goa/design"
. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("minesweeper", func() {
	Title("API for minesweeper game.")
	Description("This miscroservice provides back-end functionality for minesweeper game.")
	Scheme("http")
	Host("localhost:8080")
})

var _ = Resource("build", func() {
	BasePath("/build")

	//This endpoint is for checking microservice health
	//If microservice is up and running or is down.
	Action("BuildNumber", func() {
		Description("Get a welcome message")
		Routing(GET("/"))
		Response(OK, ConfirmationMedia)
		Response(NotFound)
	})
})

// ConfirmationMedia defines the media type used to render health endpoint.
var ConfirmationMedia = MediaType("application/vnd.goa.confirmation+json", func() {
	Description("Structure to check microservice health.")
	Attributes(func() {
		Attribute("message", String, "Message to confirm operation")
		Required("message")
	})
	View("default", func() {
		Attribute("message")
	})
})