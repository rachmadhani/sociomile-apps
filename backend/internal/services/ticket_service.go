package services

import (
	"errors"
	model "sociomile-apps/internal/models"
	"sociomile-apps/internal/repositories"

	"github.com/google/uuid"
)

type TicketService struct {
	ticketRepo *repositories.TicketRepository
	convRepo   *repositories.ConversationRepository
}

func NewTicketService(
	ticketRepo *repositories.TicketRepository,
	convRepo *repositories.ConversationRepository,
) *TicketService {
	return &TicketService{
		ticketRepo: ticketRepo,
		convRepo:   convRepo,
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

	if err := s.ticketRepo.CreateTicket(ticket); err != nil {
		return nil, err
	}

	return ticket, nil
}

func (s *TicketService) UpdateTicketStatus(
	ticketID uuid.UUID,
	tenantID uuid.UUID,
	status string,
) error {
	return s.ticketRepo.UpdateStatus(ticketID, tenantID, status)
}

func (s *TicketService) ListTicket(
	tenantID uuid.UUID,
	page int,
	limit int,
) ([]model.Ticket, int64, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 50 {
		limit = 10
	}

	offset := (page - 1) * limit
	return s.ticketRepo.List(tenantID, offset, limit)
}
