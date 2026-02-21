package conversation

type AgentReplyRequest struct {
	Message string `json:"message" binding:"required"`
}
