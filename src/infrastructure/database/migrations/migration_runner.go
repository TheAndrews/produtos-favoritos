package migrations

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var files = []*gormigrate.Migration{&migration202508050602,
	&migration202508060345}

func RunMigrations(db *gorm.DB) {
	// Setup DB
	// db := &gorm.Database{}
	// dbInstance := db.GetInstance()

	m := gormigrate.New(db, gormigrate.DefaultOptions, files)
	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration run successfully")
}
