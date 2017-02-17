package main

import (
	"github.com/goadesign/goa"
	"github.com/viejodecaldas/ms-object-minesweeper/src/web/app"
)

// BuildController implements the build resource.
type BuildController struct {
	*goa.Controller
}

// NewBuildController creates a build controller.
func NewBuildController(service *goa.Service) *BuildController {
	return &BuildController{Controller: service.NewController("BuildController")}
}

// BuildNumber runs the BuildNumber action.
func (c *BuildController) BuildNumber(ctx *app.BuildNumberBuildContext) error {
	// BuildController_BuildNumber: start_implement

	res := &app.GoaConfirmation{Message: "Health check"}
	return ctx.OK(res)
	// BuildController_BuildNumber: end_implement
}
