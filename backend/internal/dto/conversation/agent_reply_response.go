package conversation

type AgentReplyResponse struct {
	ConversationID string `json:"conversation_id"`
	Status         string `json:"status"`
	AssignedAgent  string `json:"assigned_agent_id"`
}
