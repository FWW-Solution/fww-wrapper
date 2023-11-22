package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ServiceName string           `mapstructure:"service_name"`
	HttpServer  HttpServerConfig `mapstructure:"http_server"`

	Database      DatabaseConfig      `mapstructure:"database"`
	Redis         RedisConfig         `mapstructure:"redis"`
	MessageStream MessageStreamConfig `mapstructure:"message_stream"`
	OpenTelemetry OpenTelemetryConfig `mapstructure:"open_telemetry"`
	GrpcServer    GrpcServerConfig    `mapstructure:"grpc"`
	GrpcClient    GrpcClientConfig    `mapstructure:"grpc_client"`
	HttpClient    HttpClientConfig    `mapstructure:"http_client"`
	Logger        LoggerConfig        `mapstructure:"logger"`
	ObjectStorage MinioConfig         `mapstructure:"object_storage"`
	Midtrans      MidtransConfig      `mapstructure:"midtrans"`
	Email         EmailConfig         `mapstructure:"email"`
}

type EmailConfig struct {
	Server       string `mapstructure:"server"`
	SmtpPort     string `mapstructure:"smtp_port"`
	SkipSSL      bool   `mapstructure:"skip_ssl"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	EmailAddress string `mapstructure:"email_address"`
}

type HttpServerConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type MidtransConfig struct {
	IsProduction bool   `mapstructure:"is_production"`
	ServerKey    string `mapstructure:"server_key"`
	ClientKey    string `mapstructure:"client_key"`
}

type MinioConfig struct {
	Endpoint        string `mapstructure:"endpoint"`
	AccessKeyID     string `mapstructure:"access_key_id"`
	SecretAccessKey string `mapstructure:"secret_access_key"`
	UseSSL          bool   `mapstructure:"use_ssl"`
	BucketName      string `mapstructure:"bucket_name"`
}

type HttpClientConfig struct {
	Host                string  `mapstructure:"host"`
	Port                string  `mapstructure:"port"`
	Timeout             int     `mapstructure:"timeout"`
	ConsecutiveFailures int     `mapstructure:"consecutive_failures"`
	ErrorRate           float64 `mapstructure:"error_rate"` // 0.001 - 0.999
	Threshold           int     `mapstructure:"threshold"`
	Type                string  `mapstructure:"type"` // consecutive, error_rate
}

type LoggerConfig struct {
	IsVerbose       bool   `mapstructure:"is_verbose"`
	LoggerCollector string `mapstructure:"logger_collector"`
}

type GrpcServerConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type GrpcClientConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"db_name"`
	SSL          string `mapstructure:"ssl"`
	SchemaName   string `mapstructure:"schema_name"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	Timeout      int    `mapstructure:"timeout"`
}

type OpenTelemetryConfig struct {
	ExporterType             string `mapstructure:"exporter_type"`
	OtelExporterOLTPEndpoint string `mapstructure:"otel_exporter_otlp_endpoint"`
	OtelExporterOTLPInsecure bool   `mapstructure:"otel_exporter_otlp_insecure"`
}

type MessageStreamConfig struct {
	Host           string `mapstructure:"host"`
	Port           string `mapstructure:"port"`
	Username       string `mapstructure:"username"`
	Password       string `mapstructure:"password"`
	ExchangeName   string `mapstructure:"exchange_name"`
	PublishTopic   string `mapstructure:"publish_topic"`
	SubscribeTopic string `mapstructure:"subscribe_topic"`
}

type RedisConfig struct {
	Host            string        `mapstructure:"host"`
	Port            string        `mapstructure:"port"`
	Username        string        `mapstructure:"username"`
	Password        string        `mapstructure:"password"`
	DB              int           `mapstructure:"db"`
	MaxRetries      int           `mapstructure:"max_retries"`
	PoolFIFO        bool          `mapstructure:"pool_fifo"`
	PoolSize        int           `mapstructure:"pool_size"`
	PoolTimeout     time.Duration `mapstructure:"pool_timeout"`
	MinIdleConns    int           `mapstructure:"min_idle_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
}

// InitConfig reads config file from path

func InitConfig() *Config {
	// // Viper add remote provider
	// viper.AddRemoteProvider("consul", "localhost:8500", "config/hello-service.json")
	// viper.SetConfigType("json")
	// err := viper.ReadRemoteConfig()
	// if err != nil {
	// 	panic(err)
	// }

	// Viper read file from path
	viper.AddConfigPath("./internal/config")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var Cfg Config

	// Viper Populates the struct
	err = viper.Unmarshal(&Cfg)
	if err != nil {
		panic(err)
	}
	return &Cfg
}
