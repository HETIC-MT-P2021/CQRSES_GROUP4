package elasticsearch

import (
	db "github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
)

// Repository Makes process on any database
type Repository interface {
	SetUpIndexes() error

	IsClientReady(string) error
	CreateIndexIfNotExists(string) error

	StoreEvent(event db.Event) error
	LoadEvents(string) ([]db.Event, error)

	StoreReadmodel(db.Article) error
	UpdateReadmodel(string, db.Article) error
	GetReadmodel(string) (db.Article, error)
}
