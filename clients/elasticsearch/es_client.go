package elasticsearch

import (
	"context"
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
	Index(string, interface{}) (*elastic.IndexResponse, error)
}

func Init() {

	var err error
	log := logger.GetLogger()
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
	)

	if err != nil {
		panic(err)
	}

	Client.setClient(client)
}

func (c *esClientStruct) Index(index string, obj interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().Index(index).BodyJson(obj).Do(ctx)

	if err != nil {
		logger.Error("error when trying to index document in elasticSearch", err)
		return nil, err
	}
	return result, nil
}

func (c *esClientStruct) setClient(client *elastic.Client) {
	c.client = client
}
