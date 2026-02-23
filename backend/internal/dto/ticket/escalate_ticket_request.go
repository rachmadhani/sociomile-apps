package ticket

type EscalateTicketRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Priority    string `json:"priority" binding:"omitempty,oneof=low medium high"`
}
