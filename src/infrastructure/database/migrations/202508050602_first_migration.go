package migrations

import (
	"produtos-favoritos/src/domain/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var migration202508050602 = gormigrate.Migration{
	ID: "202508050602",
	Migrate: func(tx *gorm.DB) error {
		tx.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

		tx.AutoMigrate(
			&models.Customer{},
		)

		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		return nil
	},
}
