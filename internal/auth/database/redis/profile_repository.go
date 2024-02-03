package redis

import (
	"context"
	"fmt"

	"github.com/charmingruby/clize/internal/auth/domain/profile"
	rq "github.com/charmingruby/clize/pkg/database/redis"
	"github.com/charmingruby/clize/pkg/errors"
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
		return nil, &errors.ResourceNotFoundError{
			Entity:  "profile",
			Message: errors.NewResourceNotFoundErrorMessage("profile"),
		}
	}

	return profile, nil
}

func (pr *RedisProfileRepository) FindByUsername(username string) (*profile.Profile, error) {
	key := fmt.Sprintf("%s%s", profilePattern, username)

	profile, err := rq.Get[profile.Profile](*pr.rc, pr.ctx, key)

	if err != nil {
		return nil, &errors.ResourceNotFoundError{
			Entity:  "profile",
			Message: errors.NewResourceNotFoundErrorMessage("profile"),
		}
	}

	return profile, nil
}

func (pr *RedisProfileRepository) FindById(id string) (*profile.Profile, error) {
	key := fmt.Sprintf("%s%s", profilePattern, id)

	profile, err := rq.Get[profile.Profile](*pr.rc, pr.ctx, key)

	if err != nil {
		return nil, &errors.ResourceNotFoundError{
			Entity:  "profile",
			Message: errors.NewResourceNotFoundErrorMessage("profile"),
		}
	}

	return profile, nil
}
