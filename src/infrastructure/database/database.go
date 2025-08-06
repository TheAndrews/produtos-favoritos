package database

import (
	"fmt"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"produtos-favoritos/src/infrastructure/config"
)

type Database struct {
	db   *gorm.DB
	once sync.Once
}

func (r *Database) connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable TimeZone=UTC",
		config.DB_HOST,
		config.DB_PORT,
		config.DB_USER,
		config.DB_NAME,
		config.DB_PASS,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB from gorm.DB: %w", err)
	}

	// Connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func (r *Database) GetInstance() *gorm.DB {
	r.once.Do(func() {
		var err error
		r.db, err = r.connect()
		if err != nil {
			panic("failed to initialize DB: " + err.Error())
		}
	})
	return r.db
}

func (r *Database) Stop() {
	if r.db != nil {
		sqlDB, err := r.db.DB()
		if err == nil {
			_ = sqlDB.Close()
		}
	}
}
