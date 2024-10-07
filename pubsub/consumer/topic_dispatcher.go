package consumer

import (
	"demo_test_worker/mod/constants"
	"fmt"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type MessageHandler interface {
	Handle(Message) error
}
type Message struct {
	*sarama.ConsumerMessage
}

func (c Message) MarshalLogObject(e zapcore.ObjectEncoder) error {
	if c.ConsumerMessage == nil {
		return nil
	}
	e.AddString("Topic", c.ConsumerMessage.Topic)
	e.AddInt32("Partition", c.ConsumerMessage.Partition)
	e.AddInt64("Offset", c.ConsumerMessage.Offset)
	return nil
}

type TopicDispatcherOption struct {
	Topic   string
	Handler MessageHandler
}

type TopicDispatcher struct {
	handlersMap map[string]MessageHandler
}

func NewTopicDispatcher(opts ...TopicDispatcherOption) *TopicDispatcher {
	handlersMap := make(map[string]MessageHandler)
	for _, o := range opts {
		handlersMap[o.Topic] = o.Handler
	}
	return &TopicDispatcher{
		handlersMap: handlersMap,
	}
}

func (dispatcher *TopicDispatcher) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (dispatcher *TopicDispatcher) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (dispatcher *TopicDispatcher) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	defer func() {
		if e := recover(); e != nil {
			zap.L().Error("recovered from panic", zap.String(constants.ErrorRaw, fmt.Sprintf(" %v", e)))
		}
	}()
	topic := claim.Topic()
	partition := claim.Partition()
	handler := dispatcher.handlersMap[topic]
	for message := range claim.Messages() {
		if message == nil {
			zap.L().Error("received nil message", zap.String("topic", topic), zap.Int32("partition", partition))
			continue
		}
		msg := Message{message}
		zap.L().Info("message info",
			zap.String("topic", message.Topic),
			zap.Int32("partition", message.Partition),
			zap.Int64("offset", message.Offset),
			zap.String("value", string(message.Value)))

		_ = withRecoverError(func() error {
			return handler.Handle(msg)
		})
		session.MarkMessage(message, "success")
	}
	return nil
}

func (dispatcher *TopicDispatcher) Topics() []string {
	topics := make([]string, 0, len(dispatcher.handlersMap))
	for t := range dispatcher.handlersMap {
		topics = append(topics, t)
	}
	return topics
}

func withRecoverError(f func() error) error {
	defer func() {
		if e := recover(); e != nil {
			zap.L().Error("recovered from panic", zap.String(constants.ErrorRaw, fmt.Sprint(e)))
		}
	}()
	return f()
}
