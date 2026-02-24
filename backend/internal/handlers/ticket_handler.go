package handlers

import (
	"net/http"
	dtoTicket "sociomile-apps/internal/dto/ticket"
	"sociomile-apps/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TicketHandler struct {
	service *services.TicketService
}

func NewTicketHandler(s *services.TicketService) *TicketHandler {
	return &TicketHandler{
		service: s,
	}
}

func (h *TicketHandler) EscalateTicket(c *gin.Context) {
	var req dtoTicket.EscalateTicketRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	convID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid conversation id"})
		return
	}

	tenantID, err := uuid.Parse(c.GetString("tenant_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tenant id"})
		return
	}

	userID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	ticket, err := h.service.EscalateTicket(
		convID,
		tenantID,
		userID,
		req.Title,
		req.Description,
		req.Priority,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var assignedAgent *string
	if ticket.AssignedAgentID != nil {
		agentStr := ticket.AssignedAgentID.String()
		assignedAgent = &agentStr
	}

	c.JSON(http.StatusOK, dtoTicket.EscalateTicketResponse{
		ID:             ticket.ID.String(),
		ConversationID: ticket.ConversationID.String(),
		Status:         string(ticket.Status),
		Priority:       string(ticket.Priority),
		AssignedAgent:  assignedAgent,
		Title:          ticket.Title,
		Description:    ticket.Description,
		CreatedAt:      ticket.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:      ticket.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	})

}

func (h *TicketHandler) UpdateStatus(c *gin.Context) {
	var req dtoTicket.EscalateTicketUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticketID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ticket id"})
		return
	}

	tenantID := req.TenantID

	if err := h.service.UpdateTicketStatus(ticketID, tenantID, req.Status); err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "ticket not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ticket status updated"})
}

func (h *TicketHandler) List(c *gin.Context) {
	tickets, total, _ := h.service.ListTicket(
		c.GetInt("page"),
		c.GetInt("limit"),
	)

	resp := make([]dtoTicket.TicketListItem, 0)
	for _, t := range tickets {
		var assignedAgent *string
		if t.AssignedAgentID != nil {
			agentStr := t.AssignedAgentID.String()
			assignedAgent = &agentStr
		}

		resp = append(resp, dtoTicket.TicketListItem{
			ID:             t.ID.String(),
			ConversationID: t.ConversationID.String(),
			TenantID:       t.TenantID.String(),
			Status:         string(t.Status),
			Priority:       string(t.Priority),
			AssignedAgent:  assignedAgent,
		})
	}

	c.JSON(http.StatusOK, dtoTicket.ListTicketResponse{
		Data:  resp,
		Total: total,
		Page:  c.GetInt("page"),
		Limit: c.GetInt("limit"),
	})
}
