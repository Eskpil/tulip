package gateway

type Message struct {
	Subject  string      `json:"subject"`
	EntityId string      `json:"entity_id"`
	Payload  interface{} `json:"payload"`
}
