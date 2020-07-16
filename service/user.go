package service

import (
	"relation-graph/models"
	"relation-graph/repositories"
)

type UserService struct {
	SkuRepo repositories.SkuRepository
}


func (s *UserService) GetBoughtSku(userId int) []models.Sku {
	var (
		skuList []models.Sku
	)

	skuList = s.SkuRepo.QueryBoughtSku(userId)
	return skuList
}
