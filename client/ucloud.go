package client

import (
	"relation-graph/models"
	"strconv"
)

type UCloudClient struct {

}

var (
	ucClient *UCloudClient
)

func GetUCloudClient() *UCloudClient {
	return ucClient
}

func (uc *UCloudClient) GetUserInfo(userId int) *models.User {
	return &models.User{
		UserId:   userId,
		UserName: strconv.Itoa(userId) + "_name",
	}
}

func init() {
	// todo ucloud用户信息服务
	ucClient = &UCloudClient{}
}