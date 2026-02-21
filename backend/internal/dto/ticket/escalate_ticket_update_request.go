package ticket

type EscalateTicketUpdateRequest struct {
	Status string `json:"status" binding:"required, oneof=open closed in_progress resolved"`
}
