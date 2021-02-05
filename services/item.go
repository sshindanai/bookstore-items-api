package services

import (
	"github.com/sshindanai/repo/bookstore-items-api/domain/models"
	"github.com/sshindanai/repo/bookstore-items-api/domain/repository"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(*models.Item) <-chan *models.ItemConcurrent
	GetItemList(chan *models.ItemConcurrent)
	Get(string, chan *models.ItemConcurrent)
	Update(string, chan *models.ItemConcurrent)
	Delete(string, chan *models.ItemConcurrent)
}

type itemsService struct{}

func (s *itemsService) Create(item *models.Item) <-chan *models.ItemConcurrent {
	output := make(chan *models.ItemConcurrent)

	if err := item.Validate(); err != nil {
		result := models.ItemConcurrent{
			Error: err,
		}
		output <- &result
		return output
	}

	go repository.ItemsRepository.Create(item, output)

	return output
}

func (s *itemsService) GetItemList(output chan *models.ItemConcurrent) {

}

func (s *itemsService) Get(itemID string, output chan *models.ItemConcurrent) {

}

func (s *itemsService) Update(itemID string, output chan *models.ItemConcurrent) {

}

func (s *itemsService) Delete(itemID string, output chan *models.ItemConcurrent) {

}
