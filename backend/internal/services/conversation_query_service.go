package services

import (
	"context"
	"sociomile-apps/internal/cache"
	model "sociomile-apps/internal/models"
	"sociomile-apps/internal/repositories"

	"github.com/google/uuid"
)

type ConversationQueryService struct {
	convRepo *repositories.ConversationRepository
	cache    *cache.ConversationCache
}

func NewConversationQueryService(convRepo *repositories.ConversationRepository, cache *cache.ConversationCache) *ConversationQueryService {
	return &ConversationQueryService{convRepo: convRepo, cache: cache}
}

func (s *ConversationQueryService) List(
	ctx context.Context,
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

	if data, total, ok := s.cache.Get(
		ctx,
		tenantID,
		status,
		assignedAgentID,
		page,
		limit,
	); ok {
		return data, total, nil
	}

	offset := (page - 1) * limit

	data, total, err := s.convRepo.List(
		tenantID,
		status,
		assignedAgentID,
		offset,
		limit,
	)

	if err != nil {
		return nil, 0, err
	}

	if err := s.cache.Set(
		ctx,
		data,
		total,
		page,
		limit,
		tenantID,
		status,
		assignedAgentID,
	); err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

func (s *ConversationQueryService) Detail(
	id uuid.UUID,
	tenantID uuid.UUID,
) (*model.Conversation, error) {
	return s.convRepo.GetDetail(id, tenantID)
}
