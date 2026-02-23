package ticket

import "github.com/google/uuid"

type EscalateTicketUpdateRequest struct {
	TenantID uuid.UUID `json:"tenant_id" binding:"required"`
	Status   string    `json:"status" binding:"required,oneof=open closed in_progress resolved"`
}
