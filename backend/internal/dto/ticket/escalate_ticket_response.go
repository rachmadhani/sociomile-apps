package ticket

type EscalateTicketResponse struct {
	ID             string  `json:"id"`
	ConversationID string  `json:"conversation_id"`
	Status         string  `json:"status"`
	Priority       string  `json:"priority"`
	AssignedAgent  *string `json:"assigned_agent_id"`
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}
