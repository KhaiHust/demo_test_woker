package event

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type ApplicationEvent struct {
	Id          string `json:"id"`
	Event       string `json:"event"`
	ServiceCode string `json:"service_code"`
	Timestamp   int64  `json:"timestamp"`
	EventTime   int64  `json:"event_time"`
}

func NewApplicationEvent(eventName string, serviceCode string) *ApplicationEvent {
	id := uuid.New().String()
	eventTime := time.Now().UnixNano() / int64(time.Millisecond)
	return &ApplicationEvent{
		Id:          id,
		Event:       eventName,
		ServiceCode: serviceCode,
		Timestamp:   eventTime,
		EventTime:   eventTime,
	}
}
func (a ApplicationEvent) Identifier() string {
	return a.Id
}

func (a ApplicationEvent) Name() string {
	return a.Event
}

func (a ApplicationEvent) Payload() interface{} {
	return nil
}

func (a ApplicationEvent) String() (string, error) {
	return a.ToString(a)
}

func (a ApplicationEvent) ToString(obj interface{}) (string, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
