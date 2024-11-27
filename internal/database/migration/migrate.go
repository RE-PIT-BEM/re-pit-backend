package main

import (
	"fmt"
	"log"

	"github.com/RE-PIT-BEM/re-pit-backend/internal/database"
	"github.com/RE-PIT-BEM/re-pit-backend/internal/model/domain"
)

func main() {
	db, err := database.NewDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}

	tables := []interface{}{
		&domain.Request{},
	}

	fmt.Println("Dropping tables...")
	db.Migrator().DropTable(tables...)
	fmt.Println("Tables Droped")

	fmt.Println("Migrating tables...")
	db.AutoMigrate(tables...)
	fmt.Println("Tables Migrated")
}
