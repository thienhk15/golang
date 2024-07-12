package utils

import (
	"time"
)

var AppConfig *Config

type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Mode string `yaml:"mode"`
	} `yaml:"server"`
	Database PgDatabase `yaml:"database"`
	MongoDb  MongoDb    `yaml:"mongodb"`
	Kafka    struct {
		Producer KafkaProducerConfig `yaml:"producer"`
		Consumer KafkaConsumerConfig `yaml:"consumer"`
	} `yaml:"kafka"`
	Redis      RedisConfig `yaml:"redis"`
	Cors       CorsConfig  `yaml:"cors"`
	AuthConfig Auth        `yaml:"auth"`
}

type PgDatabase struct {
	IsUse        bool          `yaml:"is_use"`
	Host         string        `yaml:"host"`
	Port         int           `yaml:"port"`
	User         string        `yaml:"user"`
	Password     string        `yaml:"password"`
	DbName       string        `yaml:"dbname"`
	SslMode      string        `yaml:"sslmode"`
	MaxOpenConn  int           `yaml:"max-open-conn"`
	MaxIdleConn  int           `yaml:"max-idle-conn"`
	Timeout      time.Duration `yaml:"timeout"`
	PublicSchema string        `yaml:"public-schema"`
	ReportSchema string        `yaml:"report-schema"`
}

type MongoDb struct {
	Url             string `yaml:"url"`
	DatabaseName    string `yaml:"database_name"`
	MinPoolSize     int    `yaml:"min_pool_size"`
	MaxPoolSize     int    `yaml:"max_pool_size"`
	MaxConnIdleTime int    `yaml:"max_conn_idle_time"`
}

type KafkaProducerConfig struct {
	IsUse            bool       `yaml:"is_use"`
	BootstrapServer  string     `yaml:"bootstrap-servers"`
	SecurityProtocol string     `yaml:"security-protocol"`
	SaslMechanisms   string     `yaml:"sasl-mechanisms"`
	SaslUsername     string     `yaml:"sasl-username"`
	SaslPassword     string     `yaml:"sasl-password"`
	Acks             string     `yaml:"acks"`
	FooTopic         KafkaTopic `yaml:"topics-foo1"`
}

type KafkaConsumerConfig struct {
	IsUse            bool       `yaml:"is_use"`
	BootstrapServer  string     `yaml:"bootstrap-servers"`
	SecurityProtocol string     `yaml:"security-protocol"`
	SaslMechanisms   string     `yaml:"sasl-mechanisms"`
	SaslUsername     string     `yaml:"sasl-username"`
	SaslPassword     string     `yaml:"sasl-password"`
	GroupId          string     `yaml:"group_id"`
	FooTopic         KafkaTopic `yaml:"topics-foo1"`
}

type KafkaTopic struct {
	Name      string `yaml:"name"`
	Partition int    `yaml:"partition"`
}

type RedisConfig struct {
	IsEnabled   bool   `json:"enabled" yaml:"enabled"`
	Address     string `yaml:"address"`
	Password    string `yaml:"password"`
	MinIdleConn int    `yaml:"min-idle-conn"`
	MaxOpenConn int    `yaml:"max-open-conn"`
}

// CorsConfig represents the CORS (Cross-Origin Resource Sharing) configuration.
type CorsConfig struct {
	IsEnabled        bool     `json:"enabled" yaml:"enabled"`
	AllowedOrigins   []string `json:"allowed_origins" yaml:"allowed-origins"`
	AllowedMethods   []string `json:"allowed_methods" yaml:"allowed-methods"`
	AllowedHeaders   []string `json:"allowed_headers" yaml:"allowed-headers"`
	ExposedHeaders   []string `json:"exposed_headers" yaml:"exposed-headers"`
	AllowCredentials bool     `json:"allow_credentials" yaml:"allow-credentials"`
	MaxAge           int      `json:"max_age" binding:"gte=0" yaml:"max-age"`
}

type Auth struct {
	IsEnabled bool   `yaml:"enabled"`
	Mode      string `yaml:"mode"`
}
