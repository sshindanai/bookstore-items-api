package services

import (
	"github.com/sshindanai/repo/bookstore-items-api/domain/models"
	"github.com/sshindanai/repo/bookstore-items-api/domain/repository"
)

type ItemsServiceInterface interface {
	Create(*models.Item, chan models.ItemConcurrent)
	GetItemList(chan models.ItemConcurrent)
	Get(string, chan models.ItemConcurrent)
	Update(string, chan models.ItemConcurrent)
	Delete(string, chan models.ItemConcurrent)
}

type itemsService struct {
	Repository repository.ItemsRepositoryInterface
}

func NewItemService() ItemsServiceInterface {
	return &itemsService{
		Repository: repository.NewItemRepository(),
	}
}

func (s *itemsService) Create(item *models.Item, output chan models.ItemConcurrent) {

}

func (s *itemsService) GetItemList(output chan models.ItemConcurrent) {

}

func (s *itemsService) Get(itemID string, output chan models.ItemConcurrent) {

}

func (s *itemsService) Update(itemID string, output chan models.ItemConcurrent) {

}

func (s *itemsService) Delete(itemID string, output chan models.ItemConcurrent) {

}
