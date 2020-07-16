package controllers

import (
	"github.com/kataras/iris"
	"relation-graph/library"
)

type BoughtController struct {
	BaseController
}

func (c *BoughtController) PostCreate() library.RelationResponse {
	return library.BuildRelationResponse(0, "success", nil)
}

func init() {
	RegisterController("relation/bought", func() ([]iris.Handler, interface{}) {
		return []iris.Handler{}, &BoughtController{}
	})
}
