package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ConversationStatus string

const (
	ConversationStatusOpen     ConversationStatus = "open"
	ConversationStatusClosed   ConversationStatus = "closed"
	ConversationStatusAssigned ConversationStatus = "assigned"
)

type Conversation struct {
	ID                 uuid.UUID          `gorm:"type:char(36);primary_key"`
	TenantID           uuid.UUID          `gorm:"type:char(36);index;not null"`
	CustomerExternalID string             `gorm:"type:varchar(255);not null"`
	Status             ConversationStatus `gorm:"type:varchar(255);default:'open';not null"`
	AssignedAgentID    *string            `gorm:"type:varchar(36);index"`
	CreatedAt          time.Time          `json:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at"`
	DeletedAt          gorm.DeletedAt     `gorm:"index"`

	Tenant Tenant `gorm:"foreignKey:TenantID"`
	Agent  *User  `gorm:"foreignKey:AssignedAgentID"`

	Messages []Message
	Ticket   *Ticket
}

func (c *Conversation) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}
