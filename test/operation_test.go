package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"relation-graph/client"
	"testing"
	"time"
)

type relation struct {
	UserId int `json:"user_id"`
	Name   string `json:"name"`
	SkuId  int `json:"sku_id"`
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
		UserId: 3,
		Name:   "qxw",
		SkuId:  123,
	}
	dg = client.GetDgoClient()
	if pb, err = json.Marshal(r); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%s\n", pb)
	mu = &api.Mutation{
		SetJson:   pb,
		//CommitNow: true,
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2000 * time.Millisecond)
	defer cancelFunc()
	txn := dg.NewTxn()
	if res, err = txn.Mutate(ctx, mu); err != nil {
		fmt.Println(err.Error())
	}

	if err = txn.Commit(ctx); err != nil {
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
	q = `{
	data (func: eq(name, qxw)) {
		user_id
		name
		sku_id
	}
	}`
	dg = client.GetDgoClient()

	if res, err = dg.NewTxn().Query(context.TODO(), q); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%s\n", res.Json)
}

func init() {

}
