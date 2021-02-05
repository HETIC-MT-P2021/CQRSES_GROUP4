package elasticsearch

import "os"

var mapping map[string]string = map[string]string{
	indexReadModel:  readModelMapping,
	indexEventStore: eventStoreMapping,
}

var indexReadModel = os.Getenv("INDEX_READ_MODEL")

const readModelMapping = ""

var indexEventStore = os.Getenv("INDEX_EVENT_STORE")

const eventStoreMapping = ""
