package migrations

import (
	"produtos-favoritos/src/domain/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var migration202508060345 = gormigrate.Migration{
	ID: "202508060345",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(
			&models.Customer{},
			&models.Product{},
		)
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("wishlists")
	},
}
