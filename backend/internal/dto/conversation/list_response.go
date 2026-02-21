package conversation

type ConversationListItem struct {
	ID              string  `json:"id"`
	Status          string  `json:"status"`
	AssignedAgentID *string `json:"assigned_agent_id"`
	CustomerID      string  `json:"customer_external_id"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}

type ListConversationResponse struct {
	Data  []ConversationListItem `json:"data"`
	Total int64                  `json:"total"`
	Page  int                    `json:"page"`
	Limit int                    `json:"limit"`
}
