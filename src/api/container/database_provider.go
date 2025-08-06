package container

import (
	"produtos-favoritos/src/infrastructure/database"

	"gorm.io/gorm"
)

func ProvideGormDB() *gorm.DB {
	db := &database.Database{}
	return db.GetInstance()
}
