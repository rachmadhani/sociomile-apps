package ticket

type TicketListItem struct {
	ID             string  `json:"id"`
	ConversationID string  `json:"conversation_id"`
	Status         string  `json:"status"`
	Priority       string  `json:"priority"`
	AssignedAgent  *string `json:"assigned_agent_id"`
}

type ListTicketResponse struct {
	Data  []TicketListItem `json:"data"`
	Total int64            `json:"total"`
	Page  int              `json:"page"`
	Limit int              `json:"limit"`
}
