package redis

import (
	"context"
	"fmt"

	"github.com/charmingruby/clize/helpers"
	app "github.com/charmingruby/clize/internal/application/domain"
	rdb "github.com/go-redis/redis/v8"
)

const (
	applicationPattern = "@apps/"
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

	json, err := helpers.JSONSerialize(app)
	if err != nil {
		return err
	}

	_, err = ar.rc.Set(ar.ctx, key, json, 0).Result()
	if err != nil {
		return err
	}

	return nil
}

func (ar *RedisApplicationRepository) FindByKey(key string) (*app.Application, error) {
	data, err := ar.rc.Get(ar.ctx, key).Result()
	if err != nil {
		return nil, err
	}

	app, err := helpers.JSONDeserialize[app.Application]([]byte(data))
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (ar *RedisApplicationRepository) FindByName(name string) (*app.Application, error) {
	key := fmt.Sprintf("%s%s", applicationPattern, name)

	data, err := ar.rc.Get(ar.ctx, key).Result()
	if err != nil {
		return nil, err
	}

	app, err := helpers.JSONDeserialize[app.Application]([]byte(data))
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (ar *RedisApplicationRepository) Fetch() ([]*app.Application, error) {
	matchString := fmt.Sprintf("%s*", applicationPattern)

	iter := ar.rc.Scan(ar.ctx, 0, matchString, 0).Iterator()

	var apps []*app.Application

	for iter.Next(ar.ctx) {
		key := iter.Val()

		app, _ := ar.FindByKey(key)

		apps = append(apps, app)
	}

	if err := iter.Err(); err != nil {
		return nil, iter.Err()
	}

	return apps, nil
}

func (ar *RedisApplicationRepository) Delete(name string) error {
	key := fmt.Sprintf("%s%s", applicationPattern, name)

	_, err := ar.rc.Del(ar.ctx, key).Result()
	if err != nil {
		return err
	}

	return nil
}
