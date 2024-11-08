package database

import (
	"github.com/RE-PIT-BEM/re-pit-backend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection() (*gorm.DB, error) {
	conf := config.Load()

	postgresConfig := postgres.Config{
		DSN:                  conf.GetString("connection.supabase.uri"),
		PreferSimpleProtocol: true,
	}
	db, err := gorm.Open(postgres.New(postgresConfig), &gorm.Config{})

	return db, err
}
