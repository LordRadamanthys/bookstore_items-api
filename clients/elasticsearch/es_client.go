package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/LordRadamanthys/bookstore_utils-go/logger"
	"github.com/olivere/elastic/v7"
)

var (
	Client esClientInterface = &esClientStruct{}
)

type esClientStruct struct {
	client *elastic.Client
}

type esClientInterface interface {
	setClient(c *elastic.Client)
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string, string) (*elastic.GetResult, error)
	Search(string, elastic.Query) (*elastic.SearchResult, error)
	Delete(string, string, string) (*elastic.DeleteResponse, error)
}

func Init() {

	var err error
	log := logger.GetLogger()
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
	)

	if err != nil {
		panic(err)
	}

	Client.setClient(client)
}

func (c *esClientStruct) Index(index string, docType string, obj interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().Index(index).Type(docType).BodyJson(obj).Do(ctx)

	if err != nil {
		logger.Error("error when trying to index document in elasticSearch", err)
		return nil, err
	}
	return result, nil
}

func (c *esClientStruct) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClientStruct) Get(index string, docType string, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.client.Get().Index(index).Type(docType).Id(id).Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to get id %s", id), err)
		return nil, err
	}
	return result, nil
}

func (c *esClientStruct) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()

	result, err := c.client.Search(index).Query(query).RestTotalHitsAsInt(true).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to serach documents in index %s", index), err)
		return nil, err
	}
	return result, nil
}

func (c *esClientStruct) Delete(index string, docType string, id string) (*elastic.DeleteResponse, error) {
	ctx := context.Background()

	result, err := c.client.Delete().Index(index).Type(docType).Id(id).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to delete documents in index %s", id), err)
		return nil, err
	}
	return result, nil
}
