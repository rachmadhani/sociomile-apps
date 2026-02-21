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
	if err := r.db.Where("tenant_id = ? AND customer_external_id = ? AND status != ?", tenantID, customerExternalID, model.ConversationStatusClosed).First(&conversation).Error; err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (r *ConversationRepository) FindByIdAndTenant(
	id uuid.UUID,
	tenantID uuid.UUID,
) (*model.Conversation, error) {
	var conversation model.Conversation
	if err := r.db.Where("id = ? AND tenant_id = ?", id, tenantID).First(&conversation).Error; err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (r *ConversationRepository) List(
	tenantID uuid.UUID,
	status string,
	assignedAgentID uuid.UUID,
	offset int,
	limit int,
) ([]model.Conversation, int64, error) {
	var convs []model.Conversation
	var total int64

	query := r.db.Model(&model.Conversation{}).Where("tenant_id = ?", tenantID)

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if assignedAgentID != uuid.Nil {
		query = query.Where("assigned_agent_id = ?", assignedAgentID)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("created_at desc").Offset(offset).Limit(limit).Find(&convs).Error; err != nil {
		return nil, 0, err
	}

	return convs, total, nil
}

func (r *ConversationRepository) GetDetail(
	id uuid.UUID,
	tenantID uuid.UUID,
) (*model.Conversation, error) {
	var conversation model.Conversation
	if err := r.db.Preload("Messages", func(db *gorm.DB) *gorm.DB {
		return db.Order("messages.created_at ASC")
	}).Where("id = ? AND tenant_id = ?", id, tenantID).First(&conversation).Error; err != nil {
		return nil, err
	}
	return &conversation, nil
}

func (r *ConversationRepository) Update(conversation *model.Conversation) error {
	return r.db.Save(conversation).Error
}

func (r *ConversationRepository) Create(conversation *model.Conversation) error {
	return r.db.Create(conversation).Error
}
