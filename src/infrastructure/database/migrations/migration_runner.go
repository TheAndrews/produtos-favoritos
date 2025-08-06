package migrations

import (
	"log"

	gorm "produtos-favoritos/src/infrastructure/database"

	"github.com/go-gormigrate/gormigrate/v2"
)

var files = []*gormigrate.Migration{&migration202508050602,
	&migration202508060345}

func RunMigrations() {
	// Setup DB
	db := &gorm.Database{}
	dbInstance := db.GetInstance()

	m := gormigrate.New(dbInstance, gormigrate.DefaultOptions, files)
	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration run successfully")
}
