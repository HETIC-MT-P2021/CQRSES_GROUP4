package pkg

import (
	"net/http"

	elastic "github.com/olivere/elastic/v7"
)


func NewHandlerMock() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("success"))
	}
}

func NewElasticClientMock(url string) (*elastic.Client, error) {
	client, err := elastic.NewSimpleClient(elastic.SetURL(url))
	if err != nil {
		return nil, err
	}
	return client, nil
}