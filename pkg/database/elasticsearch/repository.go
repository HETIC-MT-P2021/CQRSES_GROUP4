package elasticsearch

import (
	db "github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
)

// Repository Makes process on any database
type Repository interface {
	Close()

	SetUpIndexes() error

	isClientReady(string) error
	createIndexIfNotExists(string) error

	StoreEvent(event db.Event) error
	LoadEvents() ([]db.Event, error)

	StoreReadmodel(db.Article) error
	UpdateReadmodel(string, db.Article) error
	GetReadmodel(string) (db.Article, error)
}

var impl Repository

func setRepository(repository Repository) {
	impl = repository
}

// Close database
func Close() {
	impl.Close()
}

// SetUpIndexes Creates indexes needed to make POST request
// index used : in mapping.go
func SetUpIndexes() error {
	return impl.SetUpIndexes()
}

// isClientReady Checks if client is ready by send packet using ping
func isClientReady(clientURL string) error {
	return impl.isClientReady(clientURL)
}

// createIndexIfNotExists in elastic database
func createIndexIfNotExists(indexName string) error {
	return impl.createIndexIfNotExists(indexName)
}

// StoreEvent in database
func StoreEvent(event db.Event) error {
	return impl.StoreEvent(event)
}

// LoadEvents from database
func LoadEvents() ([]db.Event, error) {
	return impl.LoadEvents()
}

// StoreReadmodel in database
func StoreReadmodel(article db.Article) error {
	return impl.StoreReadmodel(article)
}

// UpdateReadmodel in database
func UpdateReadmodel(aggregateArticleID string, article db.Article) error {
	return impl.UpdateReadmodel(aggregateArticleID, article)
}

// GetReadmodel from database
func GetReadmodel(aggregateID string) (db.Article, error) {
	return impl.GetReadmodel(aggregateID)
}
