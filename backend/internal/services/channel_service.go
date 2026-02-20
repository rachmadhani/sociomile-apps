package services

import (
	"errors"
	model "sociomile-apps/internal/models"
	"sociomile-apps/internal/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChannelService struct {
	convRepo *repositories.ConversationRepository
	msgRepo  *repositories.MessageRepository
}

func NewChannelService(
	convRepo *repositories.ConversationRepository,
	msgRepo *repositories.MessageRepository,
) *ChannelService {
	return &ChannelService{
		convRepo: convRepo,
		msgRepo:  msgRepo,
	}
}

func (s *ChannelService) HandleIncomingMessage(
	tenantID uuid.UUID,
	customerExternalID string,
	text string,
) (*model.Conversation, error) {
	conv, err := s.convRepo.FindOpenByCustomer(tenantID, customerExternalID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		conv = &model.Conversation{
			TenantID:           tenantID,
			CustomerExternalID: customerExternalID,
			Status:             model.ConversationStatusOpen,
		}

		if err := s.convRepo.Create(conv); err != nil {
			return nil, err
		}
	}

	msg := &model.Message{
		ConversationID: conv.ID,
		SenderType:     "customer",
		Message:        text,
	}

	if err := s.msgRepo.CreateMessage(msg); err != nil {
		return nil, err
	}

	return conv, nil

}
