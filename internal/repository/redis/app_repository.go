package redis

import (
	"github.com/charmingruby/clize/internal/domain/apps"
	rdb "github.com/go-redis/redis/v8"
)

type RedisAppRepository struct {
	rc *rdb.Client
}

func NewRedisAppRepository(rc *rdb.Client) *RedisAppRepository {
	return &RedisAppRepository{
		rc: rc,
	}
}

func (ar *RedisAppRepository) Create(app *apps.App) error {
	return nil
}
