package models

import (
	"github.com/sshindanai/bookstore-utils-go/resterrors"
)

type ItemConcurrent struct {
	Result *Item              `json:"result"`
	Error  resterrors.RestErr `json:"error"`
}

type Item struct {
	ID                string      `json:"id"`
	Seller            int64       `json:"seller"`
	Title             string      `json:"title"`
	Description       Description `json:"description"`
	Price             float32     `json:"price"`
	AvailableQuantity int         `json:"available_quantity"`
	SoldQuantity      int         `json:"sold_quantity"`
	Pictures          []Picture   `json:"pictures"`
	Video             string      `json:"video"`
	Status            string      `json:"status"`
}

type Description struct {
	PlainText string `json:"plain_text"`
	HTML      string `json:"html"`
}

type Picture struct {
	ID  int64  `json:"id"`
	URL string `json:"url"`
}

type BuyingItem struct {
	ID       string `json:"id"`
	Item     `json:"item"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
}
