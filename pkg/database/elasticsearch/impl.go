package elasticsearch

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	elastic "github.com/olivere/elastic/v7"
)

// ElasticRepository implements repository interface
type ElasticRepository struct {
	client *elastic.Client
}

// Close closes database
func (r *ElasticRepository) Close() {
}

const (
	clientURL           = "http://elasticsearch:9200"
	numberOftries       = 10
	timeToWaitInSeconds = 3
)

// MakeConnection Establish a connection with elastic client
func MakeConnection() error {
	var err error
	for index := 0; index <= numberOftries; index++ {
		es, err := newElastic(clientURL)
		if err != nil {
			time.Sleep(timeToWaitInSeconds * time.Second)
		} else {
			setRepository(es)
			break
		}
	}

	return err
}

// newElastic inits new elastic client with some default params
func newElastic(url string) (*ElasticRepository, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
	)

	if err != nil {
		return nil, err
	}

	return &ElasticRepository{client}, nil
}

// SetUpIndexes Creates needed indexes to make POST request
// @see mapping.go
func (r *ElasticRepository) SetUpIndexes() error {
	err := r.isClientReady(clientURL)
	if err != nil {
		log.Println(err)
		return err
	}

	err = r.createIndexIfNotExists(indexReadModel)
	if err != nil {
		log.Println(err)
		return err
	}

	err = r.createIndexIfNotExists(indexEventStore)
	if err != nil {
		return err
	}

	return nil
}

// isClientReady Checks if client is ready by send packet using ping
func (r *ElasticRepository) isClientReady(clientURL string) error {
	ctx := context.Background()

	var err error
	for index := 0; index <= numberOftries; index++ {
		_, _, err := r.client.Ping(clientURL).Do(ctx)
		if err != nil {
			time.Sleep(timeToWaitInSeconds * time.Second)
		} else {
			break
		}
	}

	return err
}

// createIndexIfNotExists on elasticsearch database
func (r *ElasticRepository) createIndexIfNotExists(indexName string) error {
	ctx := context.Background()

	exists, err := r.client.IndexExists(indexName).Do(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	if !exists {
		fmt.Println("key not exists")

		createIndex, err := r.client.CreateIndex(indexName).BodyString(mapping[indexName]).Do(ctx)
		if err != nil {
			return err
		}
		if !createIndex.Acknowledged {
			return errors.New("Index not acknowledged")
		}
	}

	return nil
}
