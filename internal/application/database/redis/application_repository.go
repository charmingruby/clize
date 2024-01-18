package redis

import (
	"context"
	"fmt"

	apps "github.com/charmingruby/clize/internal/application/domain"
	rdb "github.com/go-redis/redis/v8"
)

type RedisApplicationRepository struct {
	rc *rdb.Client
}

func NewRedisApplicationRepository(rc *rdb.Client) *RedisApplicationRepository {
	return &RedisApplicationRepository{
		rc: rc,
	}
}

func (ar *RedisApplicationRepository) Create(app *apps.Application) error {
	ctx := context.Background()

	identifier := fmt.Sprintf("@app/%s", app.Name)
	value := fmt.Sprintf("%s: %s", app.Name, app.Context)

	if err := ar.rc.Set(ctx, identifier, value, 0); err != nil {
		return err.Err()
	}

	return nil
}

func (ar *RedisApplicationRepository) FindByIdentifier(identifier string) *apps.Application {
	return nil
}

func (ar *RedisApplicationRepository) Delete(identifier string) error {
	return nil
}
