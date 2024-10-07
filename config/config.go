package config

import (
	"crypto/tls"
	"crypto/x509"
	"demo_test_worker/mod/constants"
	"github.com/IBM/sarama"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"io/ioutil"
	"path/filepath"
	"sync"
)

var (
	conf Config
	once sync.Once
)

type SaramaConsumerConfig struct {
	InitialOffset      string
	EnableTLS          bool
	InsecureSkipVerify bool
	ClientCertFile     string
	ClientKeyFile      string
	CAFile             string
}
type ProducerConfigs struct {
	Brokers            []string `envconfig:"BROKERS" required:"true"`
	Topic              string   `envconfig:"TOPIC" required:"true"`
	EnableTLS          bool     `envconfig:"ENABLE_TLS" default:"true"`
	InsecureSkipVerify bool     `envconfig:"INSECURE_SKIP_VERIFY" default:"false"`
	CACertFile         string   `envconfig:"CA_CERT_FILE" required:"true"`
	ClientCertFile     string   `envconfig:"CLIENT_CERT_FILE" required:"true"`
	ClientKeyFile      string   `envconfig:"CLIENT_KEY_FILE" required:"true"`
}
type ConsumerConfig struct {
	KafkaConsumerGroup string `envconfig:"KAFKA_CONSUMER_GROUP" required:"true"`
	InitialOffset      string `envconfig:"INITIAL_OFFSET" default:"newest"`
}

type ConsumerConf struct {
	Brokers       []string `envconfig:"BROKERS" required:"true"`
	ConsumerGroup string   `envconfig:"CONSUMER_GROUP" required:"true"`
	Topic         string   `envconfig:"TOPIC" required:"true"`
	InitialOffset string   `envconfig:"INITIAL_OFFSET" default:"newest"`

	EnableTLS          bool   `envconfig:"ENABLE_TLS" default:"true"`
	InsecureSkipVerify bool   `envconfig:"INSECURE_SKIP_VERIFY" default:"false"`
	CACertFile         string `envconfig:"CA_CERT_FILE" required:"true"`
	ClientCertFile     string `envconfig:"CLIENT_CERT_FILE" required:"true"`
	ClientKeyFile      string `envconfig:"CLIENT_KEY_FILE" required:"true"`
}
type Config struct {
	KafkaConfigs    ProducerConfigs
	ConsumerConfigs ConsumerConfig
}

func Load() Config {
	once.Do(initConfig)
	return conf
}
func initConfig() {
	conf = Config{}
	envconfig.MustProcess("", &conf.KafkaConfigs)
	envconfig.MustProcess("", &conf.ConsumerConfigs)
}
func NewSaramaConsumerConfig() *SaramaConsumerConfig {
	return &SaramaConsumerConfig{
		InitialOffset: "newest",
		EnableTLS:     true,
	}
}
func NewTLSConfig(insecureSkipVerify bool, clientCertFile, clientKeyFile, caCertFile string) (*tls.Config, error) {
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}

	cert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		return nil, err
	}
	tlsConfig.Certificates = []tls.Certificate{cert}

	tlsConfig.InsecureSkipVerify = insecureSkipVerify
	if !insecureSkipVerify {
		caCert, err := ioutil.ReadFile(filepath.Clean(caCertFile))
		if err != nil {
			return nil, err
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		tlsConfig.RootCAs = caCertPool
	}

	return tlsConfig, nil
}

func BuildSaramaConsumerConfig(c *SaramaConsumerConfig) *sarama.Config {
	sc := sarama.NewConfig()
	sc.Version = sarama.V1_1_0_0

	sc.Consumer.Offsets.Initial = sarama.OffsetOldest
	if c.InitialOffset == "newest" {
		sc.Consumer.Offsets.Initial = sarama.OffsetNewest
	}

	if c.EnableTLS {
		sc.Net.TLS.Enable = true
		tlsConfig, err := NewTLSConfig(
			c.InsecureSkipVerify,
			c.ClientCertFile,
			c.ClientKeyFile,
			c.CAFile,
		)
		if err != nil {
			zap.L().Panic("load tls config error", zap.String(constants.ErrorRaw, err.Error()))
		}
		sc.Net.TLS.Config = tlsConfig
	}

	return sc
}
