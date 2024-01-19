package redis

import (
	"context"
	"encoding/json"
	"fmt"

	app "github.com/charmingruby/clize/internal/application/domain"
	rdb "github.com/go-redis/redis/v8"
)

type RedisApplicationRepository struct {
	rc  *rdb.Client
	ctx context.Context
}

func NewRedisApplicationRepository(rc *rdb.Client) *RedisApplicationRepository {
	return &RedisApplicationRepository{
		rc:  rc,
		ctx: context.Background(),
	}
}

func (ar *RedisApplicationRepository) Create(app *app.Application) error {
	key := fmt.Sprintf("%s%s", applicationPattern, app.Name)

	data, err := json.Marshal(app)
	if err != nil {
		return err
	}

	_, err = ar.rc.Set(ar.ctx, key, data, 0).Result()
	if err != nil {
		return err
	}

	return nil
}

func (ar *RedisApplicationRepository) FindByName(name string) (*app.Application, error) {
	key := fmt.Sprintf("%s%s", applicationPattern, name)

	data, err := ar.rc.Get(ar.ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var app *app.Application

	if err := json.Unmarshal([]byte(data), &app); err != nil {
		return nil, err
	}

	return app, nil
}

func (ar *RedisApplicationRepository) Delete(identifier string) error {
	return nil
}
