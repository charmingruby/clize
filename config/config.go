package config

import (
	env "github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type envConfig struct {
	RedisRestUrl   string `env:"UPSTASH_REDIS_REST_URL,required"`
	RedisRestToken string `env:"UPSTASH_REDIS_REST_TOKEN,required"`
	RedisPassword  string `env:"UPSTASH_REDIS_PASSWORD,required"`
	ServerPort     string `env:"SERVER_PORT,required"`
	ServerHost     string `env:"SERVER_HOST,required"`
}

type Config struct {
	Redis  *upstashRedisCredentials
	Server *serverConfig
}

type upstashRedisCredentials struct {
	RestUrl   string
	RestToken string
	Password  string
}

type serverConfig struct {
	Port string
	Host string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	environment := envConfig{}
	err = env.Parse(&environment)
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Redis: &upstashRedisCredentials{
			RestUrl:   environment.RedisRestUrl,
			RestToken: environment.RedisRestToken,
			Password:  environment.RedisPassword,
		},

		Server: &serverConfig{
			Port: environment.ServerPort,
			Host: environment.ServerHost,
		},
	}

	return cfg, nil
}
