package test

import (
	"encoding/json"
	"fmt"
	"relation-graph/models"
	"relation-graph/service"
	"testing"
)

func Test_GetBoughtSku(t *testing.T) {
	var (
		userId int
		output []models.Sku
		res    []byte
		err    error
	)
	userId = 1
	var skuService service.UserService
	output = skuService.GetBoughtSku(userId)
	if res, err = json.Marshal(output); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(res))
}

func init() {

}
