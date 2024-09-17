package config

import (
	"github.com/IBM/sarama"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Mode              string        `yaml:"mode"`
	StatFetchURLrauzh string        `yaml:"stat_fetcher_url"`
	Postgres          PostgresFlags `yaml:"postgres"`
	Kafka             KafkaConfig   `yaml:"kafka"`
	Root              RootConfig    `yaml:"root"`
	Log               LogConfig     `yaml:"log"`
}

type PostgresFlags struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
}

type KafkaConfig struct {
	KafkaEndpoints []string       `yaml:"kafka_endpoints"`
	KafkaSettings  *sarama.Config `yaml:"kafka_settings"`
}

type RootConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type LogConfig struct {
	Level string `yaml:"level"`
}

func ParseConfig() *Config {
	b, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil
	}

	config := Config{}

	err = yaml.Unmarshal(b, &config)
	if err != nil {
		return nil
	}

	return &config
}
