package testing

import (
	config2 "demo_test_worker/mod/config"
	"demo_test_worker/mod/handler"
	"demo_test_worker/mod/pubsub"
	"demo_test_worker/mod/testing/config"
	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type TestSuite struct {
	suite.Suite
	Producer pubsub.IProducer
	Worker   *handler.DemoWorker
}

func (s *TestSuite) SetupSuite() {

	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	zap.L().Info("Test App is initializing")

	config.SetMockEnv()
	c := config2.Load()
	c.KafkaConfigs.EnableTLS = false
	var producerConfig config2.ProducerConfigs
	envconfig.MustProcess("", &producerConfig)
	producerConfig.EnableTLS = false
	producer, err := pubsub.NewProducer(producerConfig)
	if err != nil {
		zap.L().Fatal("Fail to load config producer")
		return
	}
	s.Producer = producer
	newServer := handler.NewWorker(c)
	s.Worker = newServer

	go newServer.Run()

	zap.L().Info("Test App is initialized")
}
func (s *TestSuite) TearDownSuite() {
	s.Worker.Shutdown()
	zap.L().Info("Test App is shutting down")
}
