package redis

import (
	"context"
	"fmt"

	"github.com/charmingruby/clize/internal/app/domain/application"
	"github.com/charmingruby/clize/internal/app/domain/assignment"
	rq "github.com/charmingruby/clize/pkg/database/redis"
	"github.com/charmingruby/clize/pkg/errors"
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

func (ar *RedisAssignmentRepository) CreateAndAddToApplication(applicationName string, assignment *assignment.Assignment) error {
	appKey := fmt.Sprintf("%s%s", applicationPattern, applicationName)

	app, err := rq.Get[application.Application](*ar.rc, ar.ctx, appKey)
	if err != nil {
		return &errors.ResourceNotFoundError{
			Entity:  "application",
			Message: errors.NewResourceNotFoundErrorMessage("application"),
		}
	}

	app.Assignments = append(app.Assignments, *assignment)

	app.ProgressReview()

	return rq.Create[*application.Application](*ar.rc, ar.ctx, appKey, app)
}

func (ar *RedisAssignmentRepository) Fetch() ([]*assignment.Assignment, error) {
	pattern := fmt.Sprintf("%s*", applicationPattern)

	apps, err := rq.Fetch[application.Application](*ar.rc, ar.ctx, pattern)
	if err != nil {
		return nil, err
	}

	var assignments []*assignment.Assignment

	for _, app := range apps {
		assignmentList := app.Assignments
		for _, assignment := range assignmentList {
			assignments = append(assignments, &assignment)
		}
	}

	return assignments, nil
}

func (ar *RedisAssignmentRepository) FetchByApplication(appName string) ([]assignment.Assignment, error) {
	pattern := fmt.Sprintf("%s%s", applicationPattern, appName)

	app, err := rq.Get[application.Application](*ar.rc, ar.ctx, pattern)
	if err != nil {
		return nil, err
	}

	return app.Assignments, nil
}

func (ar *RedisAssignmentRepository) FindByApplicationName(applicationName string) (*assignment.Assignment, error) {
	return nil, nil
}

func (ar *RedisAssignmentRepository) FindByTitle(title string) (*assignment.Assignment, error) {
	return nil, nil
}

func (ar *RedisAssignmentRepository) Modify(assignment *assignment.Assignment) error {
	return nil
}

func (ar *RedisAssignmentRepository) Sign(assignment *assignment.Assignment) error {
	return nil
}

func (ar *RedisAssignmentRepository) Delete(applicationName, assignmentName string) error {
	return nil
}
