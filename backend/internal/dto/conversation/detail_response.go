package conversation

type ConversationDetailResponse struct {
	ID              string            `json:"id"`
	Status          string            `json:"status"`
	AssignedAgentID *string           `json:"assigned_agent_id"`
	CustomerID      string            `json:"customer_external_id"`
	Message         []MessageResponse `json:"messages"`
	CreatedAt       string            `json:"created_at"`
	UpdatedAt       string            `json:"updated_at"`
}

type MessageResponse struct {
	SenderType string `json:"sender_type"`
	Message    string `json:"message"`
	CreatedAt  string `json:"created_at"`
}
