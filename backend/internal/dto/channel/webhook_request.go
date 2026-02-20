package channel

import "github.com/google/uuid"

type WebhookRequest struct {
	TenantID           uuid.UUID `json:"tenant_id" binding:"required,uuid"`
	CustomerExternalID string    `json:"customer_external_id" binding:"required"`
	Message            string    `json:"message" binding:"required"`
}
