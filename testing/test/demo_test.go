package test

import (
	event2 "demo_test_worker/mod/pubsub/event"
	"demo_test_worker/mod/testing"
	"github.com/stretchr/testify/suite"
	testing2 "testing"
	"time"
)

type DemoHandlerTest struct {
	testing.TestSuite
}

func (s *DemoHandlerTest) SetupTest() {
}
func TestDemoTest(t *testing2.T) {
	suite.Run(t, new(DemoHandlerTest))
}
func (s *DemoHandlerTest) TestDemoHandler_ShouldReturnSuccess() {

	event := event2.AbstractEvent{
		ApplicationEvent: event2.NewApplicationEvent("TestEvent", "TestServiceCode"),
		Payload_: event2.Payload{
			Data: "TestPayload",
		},
	}

	s.Producer.PublishAsync(&event)
	time.Sleep(5 * time.Second)

}
