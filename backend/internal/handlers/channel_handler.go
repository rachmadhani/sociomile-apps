package handlers

import (
	"net/http"
	channelDto "sociomile-apps/internal/dto/channel"
	"sociomile-apps/internal/services"

	"github.com/gin-gonic/gin"
)

type ChannelHandler struct {
	services *services.ChannelService
}

func NewChannelHandler(services *services.ChannelService) *ChannelHandler {
	return &ChannelHandler{
		services: services,
	}
}

func (h *ChannelHandler) Webhook(c *gin.Context) {
	var req channelDto.WebhookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conv, err := h.services.HandleIncomingMessage(
		req.TenantID,
		req.CustomerExternalID,
		req.Message,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, channelDto.WebhookResponse{
		ConversationID: conv.ID.String(),
		Status:         string(conv.Status),
	})
}
