package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SenderType string

const (
	SenderTypeAgent    SenderType = "agent"
	SenderTypeCustomer SenderType = "customer"
)

type Message struct {
	ID             string         `gorm:"primaryKey;type:varchar(36)"`
	ConversationID string         `gorm:"type:varchar(36);index;not null"`
	SenderType     SenderType     `gorm:"type:varchar(50);default:'agent';not null"`
	Message        string         `gorm:"type:text;not null"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`

	Conversation Conversation `gorm:"foreignKey:ConversationID"`
}

func (m *Message) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.New().String()
	return nil
}
