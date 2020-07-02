package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type MvcHandle func(*mvc.Application)
type ControllerConstructor func() ([]iris.Handler, interface{})

type BaseController struct {
	Ctx *iris.Context
}

var (
	Controllers map[string]MvcHandle
)

func RegisterController(path string, constructor ControllerConstructor) {
	var (
		handle MvcHandle
	)

	handle = func(app *mvc.Application) {
		beforeMiddleWareHandlers, controller := constructor()
		app.Router.Use(beforeMiddleWareHandlers...)
		app.Handle(controller)
	}
	if Controllers == nil {
		Controllers = make(map[string]MvcHandle)
	}
	Controllers[path] = handle
}
