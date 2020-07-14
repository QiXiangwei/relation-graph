package controllers

import (
	"github.com/kataras/iris"
	"relation-graph/library"
)

type UserController struct {
	BaseController
}

func (c *UserController) GetBoughtSku() library.RelationResponse {

	return library.BuildRelationResponse(0, "success", nil)
}

func init() {
	RegisterController("relation/user", func() ([]iris.Handler, interface{}) {
		return []iris.Handler{}, &UserController{}
	})
}