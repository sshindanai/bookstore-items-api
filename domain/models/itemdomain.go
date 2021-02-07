package models

import (
	"github.com/sshindanai/bookstore-utils-go/resterrors"
	"github.com/sshindanai/repo/bookstore-items-api/clients/elasticsearch"
)

const (
	indexItems = "items"
)

func (i *Item) Validate() resterrors.RestErr {
	if i.Title == "" {
		return resterrors.NewBadRequestError("error empty title")
	}
	return nil
}

func (i *Item) Save() resterrors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return resterrors.NewInternalServerError("error when trying to save an item", err)
	}

	// Set ID
	i.ID = result.Id

	return nil
}
