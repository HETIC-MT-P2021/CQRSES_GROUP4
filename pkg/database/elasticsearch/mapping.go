package elasticsearch

var mapping map[string]string = map[string]string{
	indexReadModel:  readModelMapping,
	indexEventStore: eventStoreMapping,
}

const indexReadModel = "read-model"
const readModelMapping = ""

const indexEventStore = "event-store"
const eventStoreMapping = ""
