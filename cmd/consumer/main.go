package main

import (
	"log"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/elasticsearch"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/rabbit"
)

func main() {
	domain.InitBusses()

	err := elasticsearch.MakeConnection()
	if err != nil {
		log.Println(err)
		return
	}

	elasticImpl := elasticsearch.NewElasticRepository(elasticsearch.ElasticClient)

	err = elasticImpl.SetUpIndexes()
	if err != nil {
		log.Println(err)
		return
	}

	err = rabbit.ConnectToRabbitMQ()
	if err != nil {
		log.Println(err)
		return
	}
	rabbit.Rabbit.Consume(domain.EventBus)
}
