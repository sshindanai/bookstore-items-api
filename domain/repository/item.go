package repository

import "github.com/sshindanai/repo/bookstore-items-api/domain/models"

type ItemsRepositoryInterface interface {
	Create(*models.Item, chan models.ItemConcurrent)
	Get(string, chan models.ItemConcurrent)
}

type itemsRepository struct{}

func NewItemRepository() ItemsRepositoryInterface {
	return &itemsRepository{}
}

func (i *itemsRepository) Create(item *models.Item, output chan models.ItemConcurrent) {

}

func (i *itemsRepository) Get(itemID string, output chan models.ItemConcurrent) {

}
