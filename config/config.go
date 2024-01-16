package config

import (
	env "github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type envConfig struct {
	RedisRestUrl      string `env:"UPSTASH_REDIS_REST_URL,required"`
	RedisRestToken    string `env:"UPSTASH_REDIS_REST_TOKEN,required"`
	Auth0Domain       string `env:"AUTH0_DOMAIN,required"`
	Auth0ClientId     string `env:"AUTH0_CLIENT_ID,required"`
	Auth0ClientSecret string `env:"AUTH0_CLIENT_SECRET,required"`
	ServerHTTPPort    int    `env:"SERVER_HTTP_PORT,required"`
	ServerHTTPHost    string `env:"SERVER_HTTP_HOST,required"`
}

type Config struct {
	Redis  *UpstashRedisCredentials
	Auth0  *Auth0Credentials
	Server *ServerHTTP
}

type UpstashRedisCredentials struct {
	RestUrl   string
	RestToken string
}

type Auth0Credentials struct {
	Domain       string
	ClientId     string
	ClientSecret string
}

type ServerHTTP struct {
	Host string
	Port int
}

func NewConfig() (*Config, error) {
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
		Redis: &UpstashRedisCredentials{
			RestUrl:   environment.RedisRestUrl,
			RestToken: environment.RedisRestToken,
		},
		Auth0: &Auth0Credentials{
			Domain:       environment.Auth0Domain,
			ClientId:     environment.Auth0ClientId,
			ClientSecret: environment.Auth0ClientSecret,
		},
		Server: &ServerHTTP{
			Host: environment.ServerHTTPHost,
			Port: environment.ServerHTTPPort,
		},
	}

	return cfg, nil
}
