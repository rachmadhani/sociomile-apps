package handlers

import (
	"net/http"
	convDTO "sociomile-apps/internal/dto/conversation"
	"sociomile-apps/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ConversationQueryHandler struct {
	service *services.ConversationQueryService
}

func NewConversationQueryHandler(s *services.ConversationQueryService) *ConversationQueryHandler {
	return &ConversationQueryHandler{service: s}
}

func (h *ConversationQueryHandler) List(c *gin.Context) {
	var req convDTO.ListConversationRequest
	_ = c.ShouldBindQuery(&req)

	tenantID := c.GetString("tenant_id")

	var assignedAgentID uuid.UUID
	if req.AssignedAgent != "" {
		parsed, err := uuid.Parse(req.AssignedAgent)
		if err == nil {
			assignedAgentID = parsed
		}
	}

	convs, total, err := h.service.List(
		c.Request.Context(),
		uuid.Must(uuid.Parse(tenantID)),
		req.Status,
		assignedAgentID,
		req.Page,
		req.Limit,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to List"})
		return
	}

	items := make([]convDTO.ConversationListItem, 0)
	for _, c := range convs {
		var assignedAgentID *string
		if c.AssignedAgentID != nil {
			idStr := c.AssignedAgentID.String()
			assignedAgentID = &idStr
		}

		items = append(items, convDTO.ConversationListItem{
			ID:              c.ID.String(),
			Status:          string(c.Status),
			AssignedAgentID: assignedAgentID,
			CustomerID:      c.CustomerExternalID,
			CreatedAt:       c.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:       c.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, convDTO.ListConversationResponse{
		Data:  items,
		Total: total,
		Page:  req.Page,
		Limit: req.Limit,
	})
}

func (h *ConversationQueryHandler) Detail(c *gin.Context) {
	id := c.Param("id")
	tenantID := c.GetString("tenant_id")

	conv, err := h.service.Detail(uuid.Must(uuid.Parse(id)), uuid.Must(uuid.Parse(tenantID)))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "conversation not found"})
		return
	}

	msgs := make([]convDTO.MessageResponse, 0)
	for _, m := range conv.Messages {
		msgs = append(msgs, convDTO.MessageResponse{
			SenderType: string(m.SenderType),
			Message:    m.Message,
			CreatedAt:  m.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	var assignedAgentID *string
	if conv.AssignedAgentID != nil {
		idStr := conv.AssignedAgentID.String()
		assignedAgentID = &idStr
	}

	c.JSON(http.StatusOK, convDTO.ConversationDetailResponse{
		ID:              conv.ID.String(),
		Status:          string(conv.Status),
		AssignedAgentID: assignedAgentID,
		CustomerID:      conv.CustomerExternalID,
		Message:         msgs,
		CreatedAt:       conv.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:       conv.UpdatedAt.Format("2006-01-02 15:04:05"),
	})
}
