package client

import (
	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

var (
	dgoClient *dgo.Dgraph
)

func GetDgoClient() *dgo.Dgraph {
	return dgoClient
}

func Init() {
	var (
		conn *grpc.ClientConn
		err  error
	)

	if conn, err = grpc.Dial("localhost:9080", grpc.WithInsecure()); err != nil {
		return
	}

	dgoClient = dgo.NewDgraphClient( api.NewDgraphClient(conn))
	return
}
