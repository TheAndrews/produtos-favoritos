package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var migration202508060560 = gormigrate.Migration{
	ID: "202508060560",
	Migrate: func(tx *gorm.DB) error {
		// Drop existing foreign key constraint (replace the constraint name with the actual one in your DB)
		if err := tx.Exec(`ALTER TABLE wishlists DROP CONSTRAINT IF EXISTS fk_wishlists_customer`).Error; err != nil {
			return err
		}

		// Add foreign key with ON DELETE CASCADE
		if err := tx.Exec(`
			ALTER TABLE wishlists
			ADD CONSTRAINT fk_wishlists_customer
			FOREIGN KEY (customer_id)
			REFERENCES customers(id)
			ON DELETE CASCADE
		`).Error; err != nil {
			return err
		}

		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		// Rollback: drop the FK with cascade and add the original FK without cascade
		if err := tx.Exec(`ALTER TABLE wishlists DROP CONSTRAINT IF EXISTS fk_wishlists_customer`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`
			ALTER TABLE wishlists
			ADD CONSTRAINT fk_wishlists_customer
			FOREIGN KEY (customer_id)
			REFERENCES customers(id)
		`).Error; err != nil {
			return err
		}
		return nil
	},
}
