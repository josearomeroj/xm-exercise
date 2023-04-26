package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"strings"
	"time"
)

var (
	v = validator.New()
)

type dbConfig struct {
	Url    string `validate:"required,url" mapstructure:"URL"`
	DbName string `validate:"required" mapstructure:"DBNAME"`
}

type authConfig struct {
	PrivateKey        string        `validate:"required" mapstructure:"PRIVATE_KEY"`
	PublicKey         string        `validate:"required" mapstructure:"PUBLIC_KEY"`
	JWTValidityMillis time.Duration `validate:"required" mapstructure:"JWT_VALIDITY_MILLIS"`
}

type kafkaConfig struct {
	KafakaUrl  string `validate:"required,url" mapstructure:"KAFKA_URL"`
	KafkaTopic string `validate:"required" mapstructure:"KAFKA_EVENT_TOPIC"`
}

type Config struct {
	Port        uint16      `validate:"required,min=1" mapstructure:"PORT"`
	DBConfig    dbConfig    `mapstructure:"DB_CONFIG"`
	AuthConfig  authConfig  `mapstructure:"AUTH_CONFIG"`
	KafkaConfig kafkaConfig `mapstructure:"KAFKA_CONFIG"`
}

// LoadConfig from yaml or from env config
func LoadConfig(path string) (*Config, error) {
	vip := viper.New()
	vip.AutomaticEnv()
	vip.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	vip.SetConfigFile(path)
	vip.SetConfigType("yaml")

	c := Config{}
	if err := vip.ReadInConfig(); err != nil {
		return nil, err
	} else if err := vip.Unmarshal(&c); err != nil {
		return nil, err
	} else if err := v.Struct(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
