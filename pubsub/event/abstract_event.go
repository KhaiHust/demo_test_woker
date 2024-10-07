package event

import "github.com/gin-gonic/gin"

type Payload struct {
	Data      interface{} `json:"data"`
	Attribute string      `json:"attribute,omitempty"`
	Tags      interface{} `json:"tags,omitempty"`
}
type AbstractEvent struct {
	*ApplicationEvent
	Payload_ Payload `json:"payload"`
}

func NewEvent(ctx *gin.Context, eventName string, serviceCode string, payloadData interface{}) *AbstractEvent {
	absEvent := AbstractEvent{
		ApplicationEvent: NewApplicationEvent(eventName, serviceCode),
	}

	absEvent.Payload_ = Payload{
		Data: payloadData,
	}
	return &absEvent
}
