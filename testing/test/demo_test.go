package test

import (
	event2 "demo_test_worker/mod/pubsub/event"
	"demo_test_worker/mod/testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	testing2 "testing"
)

type DemoHandlerTest struct {
	testing.TestSuite
}

func (s *DemoHandlerTest) SetupTest() {
	s.TestSuite.SetupTest()
}
func TestDemoTest(t *testing2.T) {
	suite.Run(t, new(DemoHandlerTest))
}
func (s *DemoHandlerTest) TestDemoHandler_ShouldReturnSuccess() {
	gin.SetMode(gin.TestMode)
	event := event2.AbstractEvent{
		ApplicationEvent: event2.NewApplicationEvent("TestEvent", "TestServiceCode"),
		Payload_: event2.Payload{
			Data: "TestPayload",
		},
	}
	s.Worker.Run()
	zap.L().Info("Event published", zap.Any("event", event))
	s.Producer.PublishAsync(&event)
}
