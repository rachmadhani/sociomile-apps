package services

import (
	"errors"
	model "sociomile-apps/internal/models"
	"sociomile-apps/internal/repositories"

	"github.com/google/uuid"
)

type ConversationService struct {
	convRepo *repositories.ConversationRepository
	msgRepo  *repositories.MessageRepository
}

func NewConversationService(
	convRepo *repositories.ConversationRepository,
	msgRepo *repositories.MessageRepository,
) *ConversationService {
	return &ConversationService{
		convRepo: convRepo,
		msgRepo:  msgRepo,
	}
}

func (s *ConversationService) AgentReply(
	conversationID uuid.UUID,
	tenantID uuid.UUID,
	agentId uuid.UUID,
	text string,
) (*model.Conversation, error) {
	conv, err := s.convRepo.FindByIdAndTenant(conversationID, tenantID)
	if err != nil {
		return nil, err
	}

	if conv.Status == "closed" {
		return nil, errors.New("conversation is closed")
	}

	if conv.AssignedAgentID == nil {
		conv.AssignedAgentID = &agentId
		conv.Status = "assigned"
		if err := s.convRepo.Update(conv); err != nil {
			return nil, err
		}
	}

	msg := &model.Message{
		ConversationID: conv.ID,
		SenderType:     "agent",
		Message:        text,
	}

	if err := s.msgRepo.CreateMessage(msg); err != nil {
		return nil, err
	}

	return conv, nil
}
