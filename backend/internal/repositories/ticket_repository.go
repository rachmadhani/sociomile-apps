package repositories

import (
	model "sociomile-apps/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) *TicketRepository {
	return &TicketRepository{db: db}
}

func (r *TicketRepository) ExistByConversation(conversationID uuid.UUID) (bool, error) {
	var count int64
	if err := r.db.Model(&model.Ticket{}).Where("conversation_id = ?", conversationID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *TicketRepository) CreateTicket(ticket *model.Ticket) error {
	return r.db.Create(ticket).Error
}

func (r *TicketRepository) UpdateStatus(
	id uuid.UUID,
	tenantID uuid.UUID,
	status string,
) error {
	return r.db.Model(&model.Ticket{}).Where("id = ? AND tenant_id = ?", id, tenantID).Update("status", status).Error
}

func (r *TicketRepository) List(
	tenantID uuid.UUID,
	offset int,
	limit int,
) ([]model.Ticket, int64, error) {
	var tickets []model.Ticket
	var total int64

	query := r.db.Model(&model.Ticket{}).Where("tenant_id = ?", tenantID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("created_at desc").Offset(offset).Limit(limit).Find(&tickets).Error; err != nil {
		return nil, 0, err
	}

	return tickets, total, nil
}
