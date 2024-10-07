package testing

import (
	config2 "demo_test_worker/mod/config"
	"demo_test_worker/mod/handler"
	"demo_test_worker/mod/pubsub"
	"demo_test_worker/mod/shutdown"
	"demo_test_worker/mod/testing/config"
	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type TestSuite struct {
	suite.Suite
	Producer pubsub.Producer
	Worker   *handler.StatusSynchronizationWorker
}

func (s *TestSuite) SetupTest() {
	defer shutdown.SigtermHandler().Wait()
	zap.L().Info("Test App is initializing")
	config.SetMockEnv()
	c := config2.Load()
	c.KafkaConfigs.EnableTLS = false
	var producerConfig config2.ProducerConfigs
	envconfig.MustProcess("", &producerConfig)
	producerConfig.EnableTLS = false
	producer, err := pubsub.NewProducer(producerConfig)
	if err != nil {
		zap.L().Error("Fail to load config producer")
		s.TearDownTest()
		return
	}
	s.Producer = producer
	s.Worker = handler.NewWorker(c)
	zap.L().Info("Test App is initialized")
}
func (s *TestSuite) TearDownTest() {

}
