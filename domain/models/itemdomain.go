package models

import (
	"github.com/sshindanai/bookstore-utils-go/resterrors"
)

func (i *Item) Validate() *resterrors.RestErr {
	return nil
}
