package handlers

import (
	"net/http"
	convDTO "sociomile-apps/internal/dto/conversation"
	"sociomile-apps/internal/services"
	"sociomile-apps/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ConversationHandler struct {
	service *services.ConversationService
}

func NewConversationHandler(s *services.ConversationService) *ConversationHandler {
	return &ConversationHandler{
		service: s,
	}
}

func (h *ConversationHandler) AgentReply(c *gin.Context) {
	var req convDTO.AgentReplyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conversationID := c.Param("id")

	agentID := c.GetString("user_id")
	tenantID := c.GetString("tenant_id")

	logger, logFile, err := utils.GenerateNewLogger("conversation_handler")
	if err == nil {
		logger.Printf("tenantID: %v", tenantID)
		logger.Printf("agentID: %v", agentID)
		logger.Printf("conversationID: %v", conversationID)
		logFile.Close()
	}

	conv, err := h.service.AgentReply(
		uuid.Must(uuid.Parse(conversationID)),
		uuid.Must(uuid.Parse(tenantID)),
		uuid.Must(uuid.Parse(agentID)),
		req.Message,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	assignedAgent := ""
	if conv.AssignedAgentID != nil {
		assignedAgent = conv.AssignedAgentID.String()
	}

	c.JSON(http.StatusOK, convDTO.AgentReplyResponse{
		ConversationID: conv.ID.String(),
		Status:         string(conv.Status),
		AssignedAgent:  assignedAgent,
	})
}
