package elasticsearch

import (
	"context"
	"time"

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
	Index(interface{}) (*elastic.IndexResponse, error)
}

func Init() {

	var err error
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		// elastic.SetErrorLog(),
		// elastic.SetInfoLog(),
	)

	if err != nil {
		panic(err)
	}

	Client.setClient(client)
}

func (c esClientStruct) Index(obj interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	return c.client.Index().Do(ctx)
}

func (c esClientStruct) setClient(client *elastic.Client) {
	c.client = client
}
