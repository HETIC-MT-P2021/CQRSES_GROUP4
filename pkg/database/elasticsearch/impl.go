package elasticsearch

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	elastic "github.com/olivere/elastic/v7"
)

const (
	clientURL           = "http://elasticsearch:9200"
	numberOftries       = 10
	timeToWaitInSeconds = 3
)

var ElasticClient *elastic.Client

type ElasticRepository struct {
	Client *elastic.Client
}

func NewElasticRepository(elasticClient *elastic.Client) *ElasticRepository {
	return &ElasticRepository{
		Client: elasticClient,
	}
}

// MakeConnection Establish a connection with elastic client
func MakeConnection() error {
	var err error
	for index := 0; index <= numberOftries; index++ {
		client, err := elastic.NewClient(
			elastic.SetURL(clientURL),
			elastic.SetSniff(false),
			elastic.SetHealthcheck(false),
		)
		if err != nil {
			time.Sleep(timeToWaitInSeconds * time.Second)
		} else {
			ElasticClient = client
			break
		}
	}

	return err
}

// Close closes database
/*func (r *ElasticRepository) Close() {
}*/

// SetUpIndexes Creates needed indexes to make POST request
// @see mapping.go
func (r *ElasticRepository) SetUpIndexes() error {
	err := r.IsClientReady(clientURL)
	if err != nil {
		log.Println(err)
		return err
	}

	// check if read-model exists on elastic
	err = r.CreateIndexIfNotExists(indexReadModel)
	if err != nil {
		log.Println(err)
		return err
	}

	// check if event-store exists on elastic
	err = r.CreateIndexIfNotExists(indexEventStore)
	if err != nil {
		return err
	}

	return nil
}

// isClientReady Checks if client is ready by send packet using ping
func (r *ElasticRepository) IsClientReady(clientURL string) error {
	ctx := context.Background()

	var err error
	for index := 0; index <= numberOftries; index++ {
		_, _, err := ElasticClient.Ping(clientURL).Do(ctx)
		if err != nil {
			time.Sleep(timeToWaitInSeconds * time.Second)
		} else {
			break
		}
	}

	return err
}

// createIndexIfNotExists on elasticsearch database
func (r *ElasticRepository) CreateIndexIfNotExists(indexName string) error {
	ctx := context.Background()

	exists, err := ElasticClient.IndexExists(indexName).Do(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	if !exists {
		fmt.Println("key not exists")

		createIndex, err := ElasticClient.CreateIndex(indexName).BodyString(mapping[indexName]).Do(ctx)
		if err != nil {
			return err
		}
		if !createIndex.Acknowledged {
			return errors.New("Index not acknowledged")
		}
	}

	return nil
}
