package controllers

import (
	"github.com/kataras/iris"
	"relation-graph/library"
	"relation-graph/models"
	"relation-graph/service"
	"strconv"
)

type UserController struct {
	BaseController
	Service service.UserService
}

func (c *UserController) GetBoughtSku() library.RelationResponse {

	var (
		params map[string]string
		userId int
		err    error
		output []models.Sku
	)

	params = c.Ctx.URLParams()
	if userId, err = strconv.Atoi(params["userId"]); err != nil {
		return library.BuildRelationResponse(-1, "error", err.Error())
	}

	output = c.Service.GetBoughtSku(userId)
	return library.BuildRelationResponse(0, "success", output)

}

func init() {
	RegisterController("relation/user", func() ([]iris.Handler, interface{}) {
		return []iris.Handler{}, &UserController{}
	})
}