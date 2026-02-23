package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ActivityLog struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	TenantID  uuid.UUID      `gorm:"type:char(36);index;not null"`
	EventType string         `gorm:"type:varchar(50);index;not null"`
	EntityID  uuid.UUID      `gorm:"type:char(36);index;not null"`
	Payload   string         `gorm:"type:text;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
