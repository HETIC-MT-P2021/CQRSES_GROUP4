package elasticsearch

import (
	db "github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
)

type Repository interface {
	Close()
	StoreEvent(event db.Event) error
	LoadEvents() ([]db.Event, error)
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Close()
}

func StoreEvent(event db.Event) error {
	return impl.StoreEvent(event)
}

func LoadEvents() ([]db.Event, error) {
	return impl.LoadEvents()
}
