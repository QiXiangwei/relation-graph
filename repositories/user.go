package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"golang.org/x/net/context"
	"relation-graph/client"
	"relation-graph/models"
	"strconv"
)

type UserRepository struct {

}

func (ur *UserRepository) GetUser(userId int) []*models.User {
	var (
		dg  *dgo.Dgraph
		v   map[string]string
		q   string
		res *api.Response
		err error
	)

	dg = client.GetDgoClient()
	v  = make(map[string]string)
	v["$userId"] = strconv.Itoa(userId)
	q = `query user($userId: int) {
		user (func: eq(user_id, $userId)) {
			user_id
			user_name
		}
	}`
	if res, err = dg.NewReadOnlyTxn().QueryWithVars(context.TODO(), q, v); err != nil {
		fmt.Println(err.Error())
	}
	type Root struct {
		User []*models.User `json:"user"`
	}
	root := &Root{}
	if err = json.Unmarshal(res.Json, &root); err != nil {
		fmt.Println(err.Error())
	}
	return root.User
}

func (ur *UserRepository) CreateUser(userId int) {
	var (
		userInfo *models.User
		uc       *client.UCloudClient
	)
	uc       = client.GetUCloudClient()
	userInfo = uc.GetUserInfo(userId)
}