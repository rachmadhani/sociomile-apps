package event

type Event struct {
	TenantID  string                 `json:"tenant_id"`
	EventType string                 `json:"event_type"`
	EntityID  string                 `json:"entity_id"`
	Payload   map[string]interface{} `json:"payload"`
}
