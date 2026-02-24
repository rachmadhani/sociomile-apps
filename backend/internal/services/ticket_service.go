package services

import (
	"context"
	"errors"
	"sociomile-apps/internal/cache"
	"sociomile-apps/internal/event"
	model "sociomile-apps/internal/models"
	"sociomile-apps/internal/repositories"

	"github.com/google/uuid"
)

type TicketService struct {
	ticketRepo *repositories.TicketRepository
	convRepo   *repositories.ConversationRepository
	dispatcher *event.Dispatcher
	cache      *cache.TicketCache
}

func NewTicketService(
	ticketRepo *repositories.TicketRepository,
	convRepo *repositories.ConversationRepository,
	dispatcher *event.Dispatcher,
	cache *cache.TicketCache,
) *TicketService {
	return &TicketService{
		ticketRepo: ticketRepo,
		convRepo:   convRepo,
		dispatcher: dispatcher,
		cache:      cache,
	}
}

func (s *TicketService) EscalateTicket(
	ConversationID uuid.UUID,
	tenantID uuid.UUID,
	agentID uuid.UUID,
	title string,
	desc string,
	priority string,
) (*model.Ticket, error) {
	conv, err := s.convRepo.FindByIdAndTenant(ConversationID, tenantID)
	if err != nil {
		return nil, err
	}

	exist, _ := s.ticketRepo.ExistByConversation(ConversationID)
	if exist {
		return nil, errors.New("ticket already exists")
	}

	ticket := &model.Ticket{
		ConversationID:  conv.ID,
		TenantID:        tenantID,
		AssignedAgentID: &agentID,
		Title:           title,
		Description:     desc,
		Priority:        model.TicketPriority(priority),
		Status:          model.TicketStatusOpen,
	}

	s.dispatcher.Dispatch(event.Event{
		TenantID:  tenantID.String(),
		EventType: "conversation.escalated",
		EntityID:  conv.ID.String(),
	})

	if err := s.ticketRepo.CreateTicket(ticket); err != nil {
		return nil, err
	}

	s.dispatcher.Dispatch(event.Event{
		TenantID:  tenantID.String(),
		EventType: "ticket.created",
		EntityID:  ticket.ID.String(),
		Payload: map[string]interface{}{
			"priority": priority,
		},
	})

	return ticket, nil
}

func (s *TicketService) UpdateTicketStatus(
	ticketID uuid.UUID,
	tenantID uuid.UUID,
	status string,
) error {
	err := s.ticketRepo.UpdateStatus(ticketID, tenantID, status)
	if err == nil {
		s.cache.InvalidateLists(context.Background())
	}
	return err
}

func (s *TicketService) ListTicket(
	page int,
	limit int,
) ([]model.Ticket, int64, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 50 {
		limit = 10
	}

	if data, total, ok := s.cache.Get(
		context.Background(),
		"open",
		uuid.Nil,
		page,
		limit,
	); ok {
		return data, total, nil
	}

	offset := (page - 1) * limit

	data, total, err := s.ticketRepo.List(offset, limit)

	if err != nil {
		return nil, 0, err
	}

	s.cache.Set(
		context.Background(),
		data,
		total,
		page,
		limit,
		"open",
		uuid.Nil,
	)

	return data, total, nil
}
