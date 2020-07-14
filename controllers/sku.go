package controllers

import (
	"github.com/kataras/iris"
	"relation-graph/library"
)

type SkuController struct {
	BaseController
}

func (c *SkuController) GetSomeBoughtSku() library.RelationResponse {
	return library.BuildRelationResponse(0, "success", nil)
}

func init() {
	RegisterController("relation/sku", func() ([]iris.Handler, interface{}) {
		return []iris.Handler{}, &SkuController{}
	})
}