package channel

type WebhookResponse struct {
	ConversationID string `json:"conversation_id"`
	Status         string `json:"status"`
}
