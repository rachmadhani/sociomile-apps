package services

import (
	model "sociomile-apps/internal/models"
	"sociomile-apps/internal/repositories"

	"github.com/google/uuid"
)

type ConversationQueryService struct {
	convRepo *repositories.ConversationRepository
}

func NewConversationQueryService(convRepo *repositories.ConversationRepository) *ConversationQueryService {
	return &ConversationQueryService{convRepo: convRepo}
}

func (s *ConversationQueryService) List(
	tenantID uuid.UUID,
	status string,
	assignedAgentID uuid.UUID,
	page int,
	limit int,
) ([]model.Conversation, int64, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 50 {
		limit = 10
	}

	offset := (page - 1) * limit

	return s.convRepo.List(tenantID, status, assignedAgentID, offset, limit)
}

func (s *ConversationQueryService) Detail(
	id uuid.UUID,
	tenantID uuid.UUID,
) (*model.Conversation, error) {
	return s.convRepo.GetDetail(id, tenantID)
}
