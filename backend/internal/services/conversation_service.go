package services

import (
	"errors"
	"sociomile-apps/internal/event"
	model "sociomile-apps/internal/models"
	"sociomile-apps/internal/repositories"

	"github.com/google/uuid"
)

type ConversationService struct {
	convRepo   *repositories.ConversationRepository
	msgRepo    *repositories.MessageRepository
	dispatcher *event.Dispatcher
}

func NewConversationService(
	convRepo *repositories.ConversationRepository,
	msgRepo *repositories.MessageRepository,
	dispatcher *event.Dispatcher,
) *ConversationService {
	return &ConversationService{
		convRepo:   convRepo,
		msgRepo:    msgRepo,
		dispatcher: dispatcher,
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

		s.dispatcher.Dispatch(event.Event{
			TenantID:  tenantID.String(),
			EventType: "conversation.assigned",
			EntityID:  conversationID.String(),
			Payload: map[string]interface{}{
				"agent_id": agentId.String(),
			},
		})
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
