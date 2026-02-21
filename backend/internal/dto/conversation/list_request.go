package conversation

type ListConversationRequest struct {
	Status        string `form:"status"`
	AssignedAgent string `form:"assigned_agent_id"`
	Page          int    `form:"page"`
	Limit         int    `form:"limit"`
}
