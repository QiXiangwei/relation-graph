package repositories

import (

	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	relationClient "relation-graph/client"
	"relation-graph/models"
)

type SkuRepository struct {

}

func (r *SkuRepository) QueryBoughtSku(userId int) []models.Sku {
	var (
		q       string
		dg      *dgo.Dgraph
		v       map[string]string
		res     *api.Response
		err     error
		)


	q = `query List($userId:string) {
		bought (func: eq(user_id, $userId)) {
			sku_id
			sku_name
		}
	}`

	v = make(map[string]string)
	v["$userId"] = string(userId)
	dg = relationClient.GetDgoClient()
	if res, err = dg.NewTxn().QueryWithVars(context.TODO(), q, v); err != nil {
		fmt.Println(err.Error())
	}

	type Root struct {
		Sku []models.Sku `json:"sku"`
	}
	root := &Root{}
	if err = json.Unmarshal(res.Json, &root); err != nil {
		fmt.Println(err.Error())
	}

	return root.Sku

}

func QueryLikeSku(skuId int) {

}
