package controllers

import (
	"github.com/kataras/iris"
	"relation-graph/library"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) GetHello() library.RelationResponse {
	return library.BuildRelationResponse(0, "success", nil)
}

func init() {
	RegisterController("relation/index", func() ([]iris.Handler, interface{}) {
		return []iris.Handler{}, &IndexController{}
	})
}
