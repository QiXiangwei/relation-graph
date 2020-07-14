package test

import (
	"context"
	"fmt"
	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"relation-graph/client"
	"testing"
)

func Test_Create(t *testing.T) {
	var (
		op  *api.Operation
		dg  *dgo.Dgraph
		err error
	)

	op = &api.Operation{}
	op.Schema = `
		name: string @index(fulltext, term) .
		user_id: int @index(int) @upsert .
		sku_id: int @index(int) @upsert .
	`
	dg = client.GetDgoClient()
	if err = dg.Alter(context.TODO(), op); err != nil {
		fmt.Println(err.Error())
	}
}

func init() {

}
