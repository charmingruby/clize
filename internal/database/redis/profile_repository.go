package redis

import (
	"context"
	"fmt"

	"github.com/charmingruby/clize/internal/domain/profile"
	"github.com/charmingruby/clize/internal/validation"
	rq "github.com/charmingruby/clize/pkg/database/redis"
	rdb "github.com/go-redis/redis/v8"
)

type RedisProfileRepository struct {
	rc  *rdb.Client
	ctx context.Context
}

func NewRedisProfileRepository(ctx context.Context, rc *rdb.Client) *RedisProfileRepository {
	return &RedisProfileRepository{
		rc:  rc,
		ctx: ctx,
	}
}

func (pr *RedisProfileRepository) Create(p *profile.Profile) error {
	key := fmt.Sprintf("%s%s", profilePattern, p.Username)
	return rq.Create[*profile.Profile](*pr.rc, pr.ctx, key, p)
}

func (pr *RedisProfileRepository) FindByEmail(email string) (*profile.Profile, error) {
	key := fmt.Sprintf("%s%s", profilePattern, email)

	profile, err := rq.Get[profile.Profile](*pr.rc, pr.ctx, key)

	if err != nil {
		return nil, &validation.ResourceNotFoundError{
			Entity: "profile",
		}
	}

	return profile, nil
}

func (pr *RedisProfileRepository) FindByUsername(username string) (*profile.Profile, error) {
	key := fmt.Sprintf("%s%s", profilePattern, username)

	profile, err := rq.Get[profile.Profile](*pr.rc, pr.ctx, key)

	if err != nil {
		return nil, &validation.ResourceNotFoundError{
			Entity: "profile",
		}
	}

	return profile, nil
}

func (pr *RedisProfileRepository) FindById(id string) (*profile.Profile, error) {
	key := fmt.Sprintf("%s%s", profilePattern, id)

	profile, err := rq.Get[profile.Profile](*pr.rc, pr.ctx, key)

	if err != nil {
		return nil, &validation.ResourceNotFoundError{
			Entity: "profile",
		}
	}

	return profile, nil
}
