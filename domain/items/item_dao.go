package items

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/LordRadamanthys/bookstore_utils-go/rest_errors"
	"github.com/bookstore_items-api/clients/elasticsearch"
)

const (
	indexItems = "items"
	typeItem   = "_doc"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, typeItem, i)

	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("data base errors"))
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	itemId := i.Id
	result, err := elasticsearch.Client.Get(indexItems, typeItem, i.Id)

	if err != nil {
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get  id %s", i.Id), err)
	}

	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get  id %s", i.Id), err)
	}
	if err := json.Unmarshal(bytes, i); err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse database response", err)
	}

	i.Id = itemId
	fmt.Println(result.Source)
	return nil
}
