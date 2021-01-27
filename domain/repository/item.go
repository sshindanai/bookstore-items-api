package repository

import (
	"github.com/sshindanai/repo/bookstore-items-api/domain/models"
)

var (
	ItemsRepository itemsRepositoryInterface = &itemsRepository{}
)

type itemsRepositoryInterface interface {
	Create(*models.Item, chan *models.ItemConcurrent)
	Get(string, chan *models.ItemConcurrent)
}

type itemsRepository struct{}

func (i *itemsRepository) Create(item *models.Item, output chan *models.ItemConcurrent) {
	go func() {
		if err := item.Save(); err != nil {
			result := models.ItemConcurrent{
				Error: err,
			}
			output <- &result
			return
		}

		result := models.ItemConcurrent{
			Result: item,
		}
		output <- &result
	}()
}

func (i *itemsRepository) Get(itemID string, output chan *models.ItemConcurrent) {

}
