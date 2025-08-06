package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}

	return
}
