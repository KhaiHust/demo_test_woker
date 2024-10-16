package event

type Payload struct {
	Data      interface{} `json:"data"`
	Attribute string      `json:"attribute,omitempty"`
	Tags      interface{} `json:"tags,omitempty"`
}
type AbstractEvent struct {
	*ApplicationEvent
	Payload_ Payload `json:"payload"`
}
