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
	Auth0CallbackUrl  string `env:"AUTH0_CALLBACK_URL,required"`
	Auth0LogoutUrl    string `env:"AUTH0_LOGOUT_URL,required"`
}

type Config struct {
	Redis *upstashRedisCredentials
	Auth0 *Auth0Credentials
}

type upstashRedisCredentials struct {
	RestUrl   string
	RestToken string
	Password  string
}

type Auth0Credentials struct {
	Domain       string
	ClientId     string
	ClientSecret string
	LogoutUrl    string
	CallbackUrl  string
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
		Auth0: &Auth0Credentials{
			Domain:       environment.Auth0Domain,
			ClientId:     environment.Auth0ClientId,
			ClientSecret: environment.Auth0ClientSecret,
			LogoutUrl:    environment.Auth0LogoutUrl,
			CallbackUrl:  environment.Auth0CallbackUrl,
		},
	}

	return cfg, nil
}
