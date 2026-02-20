package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleAgent UserRole = "agent"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	TenantID     uuid.UUID `gorm:"type:uuid;index;not null"`
	Name         string    `gorm:"type:varchar(255);not null"`
	Email        string    `gorm:"size:150;uniqueIndex;not null"`
	PasswordHash string    `gorm:"size:255;not null"`
	Role         UserRole  `gorm:"type:varchar(50);default:'agent'" json:"role"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Tenant Tenant `gorm:"foreignKey:TenantID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
