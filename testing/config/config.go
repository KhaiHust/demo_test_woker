package config

import "os"

func SetMockEnv() {

	// ProducerConfigs
	os.Setenv("BROKERS", "localhost:29092")
	os.Setenv("TOPIC", "myTopic")
	os.Setenv("ENABLE_TLS", "false")
	os.Setenv("INSECURE_SKIP_VERIFY", "false")
	os.Setenv("CA_CERT_FILE", "/path/to/ca_cert")
	os.Setenv("CLIENT_CERT_FILE", "/path/to/client_cert")
	os.Setenv("CLIENT_KEY_FILE", "/path/to/client_key")

	// ConsumerConfig
	os.Setenv("KAFKA_CONSUMER_GROUP", "consumer_group")
	os.Setenv("INITIAL_OFFSET", "newest")

	// ConsumerConf
	os.Setenv("BROKERS", "localhost:29092")
	os.Setenv("CONSUMER_GROUP", "consumer_group")
	os.Setenv("TOPIC", "myTopic")
	os.Setenv("INITIAL_OFFSET", "newest")
	os.Setenv("ENABLE_TLS", "false")
	os.Setenv("INSECURE_SKIP_VERIFY", "false")
	os.Setenv("CA_CERT_FILE", "/path/to/ca_cert")
	os.Setenv("CLIENT_CERT_FILE", "/path/to/client_cert")
	os.Setenv("CLIENT_KEY_FILE", "/path/to/client_key")

}
