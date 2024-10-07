package pubsub

import "demo_test_worker/mod/pubsub/consumer"

type Worker interface {
	Run()
}

type BaseWorker struct {
	ConsumerGroup *consumer.Group
	Handler       consumer.TopicConsumerGroupHandler
}
