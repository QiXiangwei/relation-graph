package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"relation-graph/client"
	"testing"
)

type relation struct {
	userId int `json:"user_id"`
	name   string `json:"name"`
	skuId  int `json:"sku_id"`
}

func Test_Add(t *testing.T) {
	var (
		r   *relation
		dg  *dgo.Dgraph
		pb  []byte
		err error
		mu  *api.Mutation
		res *api.Response
	)
	r = &relation{
		userId: 1,
		name:   "qxw",
		skuId:  123,
	}
	dg = client.GetDgoClient()
	if pb, err = json.Marshal(r); err != nil {
		fmt.Println(err.Error())
	}
	mu = &api.Mutation{
		SetJson:pb,
	}
	if res, err = dg.NewTxn().Mutate(context.TODO(), mu); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%s\n", res.Json)
}

func Test_Query(t *testing.T) {
	var (
		q   string
		dg  *dgo.Dgraph
		res *api.Response
		err error
	)
	q = `query all($a: string) {
	all(func: eq(name, $a)) {
		user_id
		name
		sku_id
	}
}`
	dg = client.GetDgoClient()

	if res, err = dg.NewTxn().QueryWithVars(context.TODO(), q, map[string]string{"$a": "qxw"}); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%s\n", res.Json)
}

func init() {

}
