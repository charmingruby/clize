package config

import (
	env "github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type envConfig struct {
	RedisRestUrl      string `env:"UPSTASH_REDIS_REST_URL,required"`
	RedisRestToken    string `env:"UPSTASH_REDIS_REST_TOKEN,required"`
	RedisPassword     string `env:"UPSTASH_REDIS_PASSWORD,required"`
	Auth0Domain       string `env:"AUTH0_DOMAIN,required"`
	Auth0ClientId     string `env:"AUTH0_CLIENT_ID,required"`
	Auth0ClientSecret string `env:"AUTH0_CLIENT_SECRET,required"`
	ServerHTTPPort    int    `env:"SERVER_HTTP_PORT,required"`
	ServerHTTPHost    string `env:"SERVER_HTTP_HOST,required"`
}

type Config struct {
	Redis  *upstashRedisCredentials
	Auth0  *auth0Credentials
	Server *serverHTTP
}

type upstashRedisCredentials struct {
	RestUrl   string
	RestToken string
	Password  string
}

type auth0Credentials struct {
	Domain       string
	ClientId     string
	ClientSecret string
}

type serverHTTP struct {
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
		Redis: &upstashRedisCredentials{
			RestUrl:   environment.RedisRestUrl,
			RestToken: environment.RedisRestToken,
			Password:  environment.RedisPassword,
		},
		Auth0: &auth0Credentials{
			Domain:       environment.Auth0Domain,
			ClientId:     environment.Auth0ClientId,
			ClientSecret: environment.Auth0ClientSecret,
		},
		Server: &serverHTTP{
			Host: environment.ServerHTTPHost,
			Port: environment.ServerHTTPPort,
		},
	}

	return cfg, nil
}
