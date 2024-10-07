package consumer

import (
	"context"
	"demo_test_worker/mod/constants"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type TopicConsumerGroupHandler interface {
	sarama.ConsumerGroupHandler
	Topics() []string
}
type Group struct {
	client        sarama.Client
	consumerGroup sarama.ConsumerGroup
	cancel        func()
}

func NewGroup(brokers []string, groupID string, config *sarama.Config) (*Group, error) {
	c, err := sarama.NewClient(brokers, config)
	if err != nil {
		return nil, err
	}
	g, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, err
	}
	consumerGroup := &Group{
		client:        c,
		consumerGroup: g,
	}
	return consumerGroup, nil
}
func (g *Group) Consume(handler TopicConsumerGroupHandler) error {
	var ctx context.Context
	ctx, g.cancel = context.WithCancel(context.Background())
	zap.L().Info("start listening to topics", zap.Any("topics", handler.Topics()))
	for {
		if ctx.Err() != nil {
			return nil
		}
		if err := g.consumerGroup.Consume(ctx, handler.Topics(), handler); err != nil {
			zap.L().Error("consumer group consume error", zap.String(constants.ErrorRaw, err.Error()))
		}
	}
}
func (g *Group) Close() error {
	g.cancel()
	return g.consumerGroup.Close()
}
