package main

import (
	"dbExample-goa/app"
	"github.com/goadesign/goa"
)

// BottleController implements the bottle resource.
type BottleController struct {
	*goa.Controller
}

// NewBottleController creates a bottle controller.
func NewBottleController(service *goa.Service) *BottleController {
	return &BottleController{Controller: service.NewController("BottleController")}
}

// Show runs the show action.
func (c *BottleController) Show(ctx *app.ShowBottleContext) error {
	// BottleController_Show: start_implement

	// Put your logic here

	res := &app.GoaExampleBottle{}
	return ctx.OK(res)
	// BottleController_Show: end_implement
}
