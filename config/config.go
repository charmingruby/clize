package config

import (
	env "github.com/caarlos0/env/v6"
)

type envConfig struct {
	RedisRestUrl   string `env:"UPSTASH_REDIS_REST_URL,required"`
	RedisRestToken string `env:"UPSTASH_REDIS_REST_TOKEN,required"`
	RedisPassword  string `env:"UPSTASH_REDIS_PASSWORD,required"`
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

	environment := envConfig{}
	err := env.Parse(&environment)
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
			Port: "8000",
			Host: "http://localhost:",
		},
	}

	return cfg, nil
}
