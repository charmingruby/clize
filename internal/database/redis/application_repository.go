package redis

import (
	"context"
	"fmt"

	"github.com/charmingruby/clize/internal/domain/application"
	"github.com/charmingruby/clize/internal/validation"
	rq "github.com/charmingruby/clize/pkg/database/redis"
	rdb "github.com/go-redis/redis/v8"
)

type RedisApplicationRepository struct {
	rc  *rdb.Client
	ctx context.Context
}

func NewRedisApplicationRepository(rc *rdb.Client, ctx context.Context) *RedisApplicationRepository {
	return &RedisApplicationRepository{
		rc:  rc,
		ctx: ctx,
	}
}

func (ar *RedisApplicationRepository) Create(app *application.Application) error {
	key := fmt.Sprintf("%s%s", applicationPattern, app.Name)
	return rq.Create[*application.Application](*ar.rc, ar.ctx, key, app)
}

func (ar *RedisApplicationRepository) FindByKey(key string) (*application.Application, error) {
	app, err := rq.Get[application.Application](*ar.rc, ar.ctx, key)

	if err != nil {
		return nil, &validation.ResourceNotFoundError{
			Entity: "application",
		}
	}

	return app, nil
}

func (ar *RedisApplicationRepository) FindByName(name string) (*application.Application, error) {
	key := fmt.Sprintf("%s%s", applicationPattern, name)

	app, err := rq.Get[application.Application](*ar.rc, ar.ctx, key)

	if err != nil {
		return nil, &validation.ResourceNotFoundError{
			Entity: "application",
		}
	}

	return app, nil
}

func (ar *RedisApplicationRepository) Fetch() ([]*application.Application, error) {
	pattern := fmt.Sprintf("%s*", applicationPattern)

	apps, err := rq.Fetch[application.Application](*ar.rc, ar.ctx, pattern)
	if err != nil {
		return nil, err
	}

	return apps, nil
}

func (ar *RedisApplicationRepository) Delete(name string) error {
	key := fmt.Sprintf("%s%s", applicationPattern, name)

	rq.Delete(*ar.rc, ar.ctx, key)

	return nil
}
