package config

import "os"

func SetMockEnv() {
	// AppConfigs
	os.Setenv("APP_NAME", "myApp")
	os.Setenv("APP_PORT", "9090")
	os.Setenv("APP_PATH", "/myapp")
	os.Setenv("APP_ADDR", "127.0.0.1")
	os.Setenv("UNLINK_URL", "testing")
	os.Setenv("URI_GET_TOKEN_HYDRA_USER_PROFILE", "/get_token_hydra")
	os.Setenv("KAFKA_TOPIC_USER_ATTRIBUTE_LINK_MANAGEMENT", "topic")
	// NewRelicConfigs
	os.Setenv("NEWRELIC_APPLICATION_NAME", "NewRelicApp")
	os.Setenv("NEWRELIC_LICENSE", "myLicenseKey")
	os.Setenv("NEWRELIC_ENABLED", "true")

	// ProducerConfigs
	os.Setenv("BROKERS", "localhost:29092")
	os.Setenv("TOPIC", "myTopic")
	os.Setenv("ENABLE_TLS", "false")
	os.Setenv("INSECURE_SKIP_VERIFY", "false")
	os.Setenv("CA_CERT_FILE", "/path/to/ca_cert")
	os.Setenv("CLIENT_CERT_FILE", "/path/to/client_cert")
	os.Setenv("CLIENT_KEY_FILE", "/path/to/client_key")

	// ExternalProducerConfigs
	os.Setenv("BOOTSTRAP_ADDRESS", "localhost:29092")
	os.Setenv("SASL_JAAS_USERNAME", "user")
	os.Setenv("SASL_JAAS_PASSWORD", "password")
	os.Setenv("SASL_ALGORITHM", "SCRAM-SHA-512")
	os.Setenv("SECURITY_PROTOCOL", "PLAINTEXT")
	os.Setenv("INSECURE_SKIP_VERIFY", "false")

	// MnSmsConfigs
	os.Setenv("BASE_URL_MN_SMS_SERVICE", "http://sms_service")
	os.Setenv("BASIC_AUTH_MN_SMS_SERVICE", "basic_auth")
	os.Setenv("X_CHANNEL", "channel")
	os.Setenv("URI_SEND_OTP_MN_SMS", "/send_otp")
	os.Setenv("URI_VERIFY_OTP_MN_SMS", "/verify_otp")

	// CTMSConfigs
	os.Setenv("BASE_URL_CTMS", "http://ctms_service")
	os.Setenv("URI_GET_TOKEN_CTMS", "/get_token")
	os.Setenv("URI_SUBMIT_TNC_CTMS", "/submit_tnc")
	os.Setenv("GRANT_TYPE_GET_TOKEN_CTMS", "grant_type")
	os.Setenv("CLIENT_ID_GET_TOKEN_CTMS", "client_id")
	os.Setenv("CLIENT_SECRET_GET_TOKEN_CTMS", "client_secret")
	os.Setenv("TNC_CODE_CTMS", "tnc_code")
	os.Setenv("TOUCH_POINT_ID_CTMS", "touch_point_id")
	os.Setenv("STATUS_CTMS", "true")

	// MnRealtimeConfigs
	os.Setenv("BASE_URL_MN_REALTIME", "http://realtime_service")
	os.Setenv("MODULE_ID_MN_REALTIME", "module_id")
	os.Setenv("URI_SEND_MESSAGE_MN_REALTIME", "/send_message")

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

	// NotificationConfigs
	os.Setenv("BASE_URL_NOTIFICATION", "http://notification_service")
	os.Setenv("BASIC_AUTH_NOTIFICATION", "basic_auth")
	os.Setenv("URI_SEND_NOTIFICATION", "/send_notification")
	os.Setenv("APP_ID_SEND_NOTIFICATION", "app_id")
	os.Setenv("TEMPLATE_ID_LINK_REQUEST_TCB_TO_ONE_U", "template_id_1")
	os.Setenv("TEMPLATE_ID_LINK_SUCCESS_ONE_U_TO_TCB", "template_id_2")
	os.Setenv("TEMPLATE_ID_LINK_FAIL_ONE_U_TO_TCB", "template_id_3")
	os.Setenv("TEMPLATE_ID_UNLINK_UPDATE_PHONE_ONE_U_TO_TCB", "template_id_4")
	os.Setenv("TEMPLATE_ID_UNLINK_SUCCESS_ONE_U_TO_TCB", "template_id_5")
	os.Setenv("TEMPLATE_ID_UNLINK_TCB_UNLINK", "template_id_6")
	os.Setenv("TEMPLATE_ID_UNLINK_TCB_PARTNER_REQUEST_FAILED", "template_id_7")

	// UserProfileConfigs
	os.Setenv("BASE_URL_USER_PROFILE", "http://user_profile_service")
	os.Setenv("URI_GET_SEARCH_USER_PROFILE", "/search_user")
	os.Setenv("BASE_URL_HYDRA_USER_PROFILE", "http://hydra_service")
	os.Setenv("BASIC_AUTH_HYDRA_USER_PROFILE", "basic_auth")
	os.Setenv("URI_GET_TOKEN_HYDRA", "/get_token_hydra")

	// LinkManagementIntConfigs
	os.Setenv("BASE_URL_LINK_MANAGEMENT_INT_SERVICE", "http://link_management_service")
	os.Setenv("BASIC_AUTH_LINK_MANAGEMENT_INT_SERVICE", "basic_auth")
	os.Setenv("URI_GET_LINK_INFO_LINK_MANAGEMENT", "/get_link_info")
	os.Setenv("URI_GET_LINK_STATUS_LINK_MANAGEMENT", "/get_link_status")
	os.Setenv("URI_CONFIRM_LINK_ACCOUNT_LINK_MANAGEMENT", "/confirm_link_account")
	os.Setenv("BASE_URL_OAUTH_HYDRA_SERVICE", "http://oauth_hydra_service")
	os.Setenv("BASIC_AUTH_OAUTH_HYDRA_SERVICE", "basic_auth")
	os.Setenv("URI_GET_TOKEN_OAUTH_HYDRA_SERVICE", "/get_token")
	os.Setenv("URI_REQUEST_LINK_ACCOUNT_LINK_MANAGEMENT", "/request_link_account")
	os.Setenv("URI_REQUEST_UN_LINK_ACCOUNT_LINK_MANAGEMENT", "/request_unlink_account")

	// MySQLConfigs
	os.Setenv("MYSQL_DSN", "user:password@tcp(localhost:3306)/dbname")
	os.Setenv("MYSQL_CONN_MAX_LIFE_TIME_SECOND", "0")
	os.Setenv("MYSQL_ENABLE_DEBUG", "false")
	os.Setenv("MYSQL_MAX_OPEN_CONNECTIONS", "25")
	os.Setenv("MYSQL_MAX_IDLE_CONNECTIONS", "5")

	// PostgreSQLConfigs
	os.Setenv("POSTGRESQL_DSN", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	os.Setenv("POSTGRESQL_CONN_MAX_LIFE_TIME_SECOND", "0")
	os.Setenv("POSTGRESQL_ENABLE_DEBUG", "false")
	os.Setenv("POSTGRESQL_MAX_OPEN_CONNECTIONS", "25")
	os.Setenv("POSTGRESQL_MAX_IDLE_CONNECTIONS", "5")

	// RedisConfigs
	os.Setenv("REDIS_URL", "redis://localhost:6379/0")
	os.Setenv("REDIS_MAX_RETRIES", "3")
	os.Setenv("REDIS_IDLE_TIMEOUT_SECOND", "300")

	// EncryptionConfigs
	os.Setenv("V1_ENCRYPTION_KEY", "encryption_key")

	// AuthenticationConfigs
	os.Setenv("OAUTH_INTROSPECT_URL", "http://oauth_introspect_url")

	// SaramaConsumerConfig
	os.Setenv("INITIAL_OFFSET", "newest")
	os.Setenv("ENABLE_TLS", "false")
	os.Setenv("INSECURE_SKIP_VERIFY", "false")
	os.Setenv("CLIENT_CERT_FILE", "/path/to/client_cert")
	os.Setenv("CLIENT_KEY_FILE", "/path/to/client_key")
	os.Setenv("CAFile", "/path/to/ca_cert")

	// HydraConfig
	os.Setenv("CLIENT_ID", "client_id")
	os.Setenv("CLIENT_SECRET", "client_secret")
	os.Setenv("TOKEN_URL", "http://token_url")
	os.Setenv("SCOPES", "internal.link_management")

	os.Setenv("NOTIFICATION_TEMPLATE_CONFIGS", "{\n  \"NotificationTemplates\": {\n    \"template_1\": {\n      \"template_id_link_request_tcb_to_one_u\": \"link_req_001\",\n      \"template_id_link_success_one_u_to_tcb\": \"link_success_001\",\n      \"template_id_link_fail_one_u_to_tcb\": \"link_fail_001\",\n      \"template_id_unlink_update_phone_one_u_to_tcb\": \"unlink_update_phone_001\",\n      \"template_id_unlink_success_one_u_to_tcb\": \"unlink_success_001\",\n      \"template_id_unlink_tcb_unlink\": \"unlink_tcb_001\",\n      \"template_id_unlink_tcb_partner_request_failed\": \"unlink_tcb_partner_fail_001\"\n    },\n    \"template_2\": {\n      \"template_id_link_request_tcb_to_one_u\": \"link_req_002\",\n      \"template_id_link_success_one_u_to_tcb\": \"link_success_002\",\n      \"template_id_link_fail_one_u_to_tcb\": \"link_fail_002\",\n      \"template_id_unlink_update_phone_one_u_to_tcb\": \"unlink_update_phone_002\",\n      \"template_id_unlink_success_one_u_to_tcb\": \"unlink_success_002\",\n      \"template_id_unlink_tcb_unlink\": \"unlink_tcb_002\",\n      \"template_id_unlink_tcb_partner_request_failed\": \"unlink_tcb_partner_fail_002\"\n    }\n  }\n}\n")
	os.Setenv("BASE_URL_USER_ATTRIBUTES", "https://api.example.com")
	os.Setenv("BASIC_AUTH_USER_ATTRIBUTES", "Basic dXNlcm5hbWU6cGFzc3dvcmQ=")
	os.Setenv("GET_USER_ATTRIBUTE_PATH", "/v1/user/attributes")
}
