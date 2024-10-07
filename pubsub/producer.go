package pubsub

import (
	"crypto/tls"
	"crypto/x509"
	"demo_test_worker/mod/config"
	"demo_test_worker/mod/pubsub/event"
	"fmt"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"time"
)

type Producer interface {
	PublishAsync(event event.Event) error
}

// compiler type check
var _ Producer = (*producer)(nil)

type producer struct {
	AsyncProducer sarama.AsyncProducer
	topic         string
}

func NewProducer(props config.ProducerConfigs) (Producer, error) {
	conf, err := BuildSaramaConfig(props)
	if err != nil {
		return nil, err
	}
	asyncProducer, err := sarama.NewAsyncProducer(props.Brokers, conf)
	if err != nil {
		return nil, err
	}
	service := &producer{
		AsyncProducer: asyncProducer,
		topic:         props.Topic,
	}
	go func() {
		for {
			select {
			case response := <-service.AsyncProducer.Successes():
				zap.L().Info("Publish message success",
					zap.String("topic", response.Topic),
					zap.String("response", fmt.Sprintf("%v", response.Value)),
					zap.Int32("partition", response.Partition),
					zap.Int64("offset", response.Offset),
				)
			case err = <-service.AsyncProducer.Errors():
				zap.L().Error("Publish message error",
					zap.String("topic", service.topic),
					zap.Error(err))
			}
		}
	}()
	return service, nil
}

func NewAsyncProducer(topic string, addr []string, config *sarama.Config) (Producer, error) {
	asyncProducer, err := sarama.NewAsyncProducer(addr, config)
	if err != nil {
		return nil, err
	}
	service := &producer{
		AsyncProducer: asyncProducer,
		topic:         topic,
	}
	go func() {
		for {
			select {
			case response := <-service.AsyncProducer.Successes():
				zap.L().Info("Publish message success",
					zap.String("topic", response.Topic),
					zap.String("response", fmt.Sprintf("%v", response.Value)),
					zap.Int32("partition", response.Partition),
					zap.Int64("offset", response.Offset),
				)
			case err = <-service.AsyncProducer.Errors():
				zap.L().Error("Publish message error",
					zap.String("topic", service.topic),
					zap.Error(err))
			}
		}
	}()

	return service, nil
}

func (k producer) PublishAsync(e event.Event) error {
	message, err := e.String()
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: k.topic,
		Key:   sarama.StringEncoder(e.Name()),
		Value: sarama.StringEncoder(message),
	}

	go func(msg *sarama.ProducerMessage) {
		k.AsyncProducer.Input() <- msg
	}(msg)
	return nil
}

func BuildSaramaConfig(config config.ProducerConfigs) (*sarama.Config, error) {
	samaraConfig := sarama.NewConfig()
	samaraConfig.Version = sarama.V1_1_0_0
	samaraConfig.Producer.Partitioner = sarama.NewHashPartitioner
	samaraConfig.Producer.RequiredAcks = sarama.WaitForAll
	samaraConfig.Producer.Return.Successes = true
	samaraConfig.Producer.Return.Errors = true
	samaraConfig.Producer.Flush.Frequency = 1 * time.Second
	//samaraConfig.Producer.Retry.Max = config.MaxRetry

	if config.EnableTLS {
		samaraConfig.Net.TLS.Enable = config.EnableTLS
		tlsConfig, err := newTLSConfig(
			config.InsecureSkipVerify,
			config.ClientCertFile,
			config.ClientKeyFile,
			config.CACertFile,
		)
		if err != nil {
			return nil, err
		}

		samaraConfig.Net.TLS.Config = tlsConfig
	}

	return samaraConfig, nil
}

func newTLSConfig(insecureSkipVerify bool, clientCertFile, clientKeyFile, caCertFile string) (*tls.Config, error) {
	tlsConfig := &tls.Config{}
	// Load client cert
	cert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		return nil, err
	}
	tlsConfig.Certificates = []tls.Certificate{cert}

	// Load CA cert
	tlsConfig.InsecureSkipVerify = insecureSkipVerify
	if !insecureSkipVerify {
		caCert, err := os.ReadFile(filepath.Clean(caCertFile))
		if err != nil {
			return nil, err
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		tlsConfig.RootCAs = caCertPool
	}

	return tlsConfig, nil
}
