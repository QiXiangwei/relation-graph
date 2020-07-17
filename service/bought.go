package service

import (
	"relation-graph/models"
	"relation-graph/repositories"
)

type BoughtService struct {
	SkuRepo  repositories.SkuRepository
	UserRepo repositories.UserRepository
}

func (bs *BoughtService) CreateBoughtRelation(userId int, skuId int) {
	var (
		userInfo []*models.User
	)

	userInfo = bs.UserRepo.GetUser(userId)
	if len(userInfo) > 0 {
		return
	}

}

func (bs *BoughtService) GetBoughtRelation(userId int, skuId int) {

}
