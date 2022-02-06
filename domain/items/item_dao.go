package items

import (
	"errors"

	"github.com/LordRadamanthys/bookstore_utils-go/rest_errors"
	"github.com/bookstore_items-api/clients/elasticsearch"
)

const (
	indexItems = "items"
)

func (i *Item) Save() *rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)

	if err != nil {
		return rest_errors.InternalServerError("error when trying to save item", errors.New("data base errors"))
	}
	i.Id = result.Id
	return nil
}
