package repositories

import (
	model "sociomile-apps/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ConversationRepository struct {
	db *gorm.DB
}

func NewConversationRepository(db *gorm.DB) *ConversationRepository {
	return &ConversationRepository{db: db}
}

func (r *ConversationRepository) FindOpenByCustomer(
	tenantID uuid.UUID,
	customerExternalID string,
) (*model.Conversation, error) {
	var conversation model.Conversation
	if err := r.db.Where("tenant_id = ? AND customer_external_id = ? AND status != ?", tenantID, customerExternalID, model.ConversationStatusOpen).First(&conversation).Error; err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (r *ConversationRepository) Create(conversation *model.Conversation) error {
	return r.db.Create(conversation).Error
}
