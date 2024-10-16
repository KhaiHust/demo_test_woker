package testing

import (
	config2 "demo_test_worker/mod/config"
	"demo_test_worker/mod/handler"
	"demo_test_worker/mod/pubsub"
	"demo_test_worker/mod/testing/config"
	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"os"
	"os/signal"
)

type TestSuite struct {
	suite.Suite
	Producer pubsub.Producer
	Worker   *handler.StatusSynchronizationWorker
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
	serverReady := make(chan bool)
	go func() {
		s.Worker = newServer
		go newServer.Run()
		if s.Worker != nil {
			serverReady <- true
		}
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit
	}()
	<-serverReady
	zap.L().Info("Test App is initialized")
}
func (s *TestSuite) TearDownSuite() {
	s.Worker.Shutdown()
	zap.L().Info("Test App is shutting down")
}
