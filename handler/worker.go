package handler

import (
	"demo_test_worker/mod/config"
	"demo_test_worker/mod/constants"
	"demo_test_worker/mod/pubsub"
	"demo_test_worker/mod/pubsub/consumer"
	event2 "demo_test_worker/mod/pubsub/event"
	"demo_test_worker/mod/shutdown"
	"encoding/json"
	"go.uber.org/zap"
)

type DemoConsumer struct {
}
type DemoWorker struct {
	pubsub.BaseWorker
}

func (s *DemoConsumer) Handle(message consumer.Message) error {
	var event event2.AbstractEvent
	err := json.Unmarshal(message.Value, &event)
	if err != nil {
		zap.L().Error("unmarshal message failed", zap.String("error", err.Error()))
		return err
	}
	zap.L().Info("event", zap.Any("event", event))
	return nil
}

func NewStatusSynchronizationConsumer() consumer.MessageHandler {
	return &DemoConsumer{}
}
func NewWorker(cfg config.Config) *DemoWorker {
	handler := NewStatusSynchronizationConsumer()
	return &DemoWorker{
		BaseWorker: pubsub.BaseWorker{
			Handler: consumer.NewTopicDispatcher(consumer.TopicDispatcherOption{
				Topic:   cfg.KafkaConfigs.Topic,
				Handler: handler,
			}),
			ConsumerGroup: initConsumerGroup(cfg),
		},
	}
}
func initConsumerGroup(conf config.Config) *consumer.Group {
	saramaConsumerConfig := config.NewSaramaConsumerConfig()
	saramaConsumerConfig.EnableTLS = conf.KafkaConfigs.EnableTLS
	saramaConsumerConfig.InsecureSkipVerify = conf.KafkaConfigs.InsecureSkipVerify
	saramaConsumerConfig.ClientCertFile = conf.KafkaConfigs.ClientCertFile
	saramaConsumerConfig.ClientKeyFile = conf.KafkaConfigs.ClientKeyFile
	saramaConsumerConfig.CAFile = conf.KafkaConfigs.CACertFile

	cg, err := consumer.NewGroup(conf.KafkaConfigs.Brokers, conf.ConsumerConfigs.KafkaConsumerGroup, config.BuildSaramaConsumerConfig(saramaConsumerConfig))
	if err != nil {
		zap.L().Panic("init consumer group error", zap.Error(err))
	}
	return cg
}
func (w *DemoWorker) Run() {
	shutdown.SigtermHandler().RegisterErrorFunc(w.ConsumerGroup.Close)
	if err := w.ConsumerGroup.Consume(w.Handler); err != nil {
		zap.L().Panic("consumer group error", zap.String(constants.ErrorRaw, err.Error()))
	}
}
func (w *DemoWorker) Shutdown() {
	zap.L().Info("shutting down worker")
	w.ConsumerGroup.Close()

}
