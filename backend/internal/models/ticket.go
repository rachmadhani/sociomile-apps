package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketStatus string
type TicketPriority string

const (
	TicketStatusOpen       TicketStatus = "open"
	TicketStatusClosed     TicketStatus = "closed"
	TicketStatusInProgress TicketStatus = "in_progress"
	TicketStatusResolved   TicketStatus = "resolved"
)

const (
	TicketPriorityLow    TicketPriority = "low"
	TicketPriorityMedium TicketPriority = "medium"
	TicketPriorityHigh   TicketPriority = "high"
)

type Ticket struct {
	ID              uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	TenantID        uuid.UUID      `gorm:"type:uuid;index;not null"`
	ConversationID  uuid.UUID      `gorm:"type:uuid;index;not null"`
	Title           string         `gorm:"type:varchar(255);not null"`
	Description     string         `gorm:"type:text;not null"`
	Status          TicketStatus   `gorm:"type:varchar(50);default:'open';not null"`
	Priority        TicketPriority `gorm:"type:varchar(50);default:'low';not null"`
	AssignedAgentID *string        `gorm:"type:varchar(36);index;not null"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Tenant       Tenant       `gorm:"foreignKey:TenantID"`
	Conversation Conversation `gorm:"foreignKey:ConversationID"`
	Agent        *User        `gorm:"foreignKey:AssignedAgentID"`
}

func (t *Ticket) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
}
