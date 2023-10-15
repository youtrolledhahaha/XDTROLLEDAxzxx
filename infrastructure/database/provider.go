package database

import (
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxentities"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxinternal"
	"github.com/youtrolledhahaha/XDTROLLEDAxzxxinternal/environment"
	"gorm.io/gorm"
	"log"
)

const tablePrefix = "v1_0_"

type Provider struct {
	Conn *gorm.DB
}

func NewProvider(configuration environment.Database) (*Provider, error) {
	switch {
	case configuration.Sqlite.IsValid():
		log.Println("Starting sqlite database")
		return NewSqliteClient(configuration.Sqlite)
	case configuration.Postgres.IsValid():
		log.Println("Starting postgres database")
		return NewPostgresClient(configuration.Postgres)
	default:
		return nil, internal.ErrNoDatabaseProvided
	}
}

func (p *Provider) Migrate() error {
	return p.Conn.AutoMigrate(
		&entities.User{},
		&entities.Device{},
		&entities.Auth{},
	)
}
