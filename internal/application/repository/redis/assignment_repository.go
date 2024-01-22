package redis

import (
	"context"

	"github.com/charmingruby/clize/internal/application/domain"
	rdb "github.com/go-redis/redis/v8"
)

type RedisAssignmentRepository struct {
	rc  *rdb.Client
	ctx context.Context
}

func NewRedisAssignmentRepository(rc *rdb.Client, ctx context.Context) *RedisAssignmentRepository {
	return &RedisAssignmentRepository{
		rc:  rc,
		ctx: ctx,
	}
}

func (ar *RedisAssignmentRepository) CreateAndAddToApplication(applicationName string, assignment *domain.Assignment) error {
	return nil
}

func (ar *RedisAssignmentRepository) FindByApplicationName(applicationName string) (*domain.Assignment, error) {
	return nil, nil
}

func (ar *RedisAssignmentRepository) FindByTitle(title string) (*domain.Assignment, error) {
	return nil, nil
}

func (ar *RedisAssignmentRepository) Modify(assignment *domain.Assignment) error {
	return nil
}

func (ar *RedisAssignmentRepository) Sign(assignment *domain.Assignment) error {
	return nil
}

func (ar *RedisAssignmentRepository) Delete(assignment *domain.Assignment) error {
	return nil
}
