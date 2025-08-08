package migrations

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var files = []*gormigrate.Migration{
	&migration202508050602,
	&migration202508060345,
	&migration202508060560}

func RunMigrations(db *gorm.DB) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, files)
	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration run successfully")
}
