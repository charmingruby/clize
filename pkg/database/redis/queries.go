package redis

import (
	"context"

	"github.com/charmingruby/clize/helpers"
	"github.com/charmingruby/clize/pkg/errors"
	rdb "github.com/go-redis/redis/v8"
)

func Create[T any](rc rdb.Client, ctx context.Context, key string, v T) error {
	json, err := helpers.JSONSerialize(v)
	if err != nil {
		return err
	}

	_, err = rc.Set(ctx, key, json, 0).Result()
	if err != nil {
		return err
	}

	return nil
}

func Get[T any](rc rdb.Client, ctx context.Context, key string) (*T, error) {
	data, err := rc.Get(ctx, key).Result()

	if err != nil {
		return nil, &errors.ResourceNotFoundError{
			Entity:  "key",
			Message: errors.NewResourceNotFoundErrorMessage(key),
		}
	}

	item, err := helpers.JSONDeserialize[T]([]byte(data))
	if err != nil {
		return nil, err
	}

	return item, nil
}

func Fetch[T any](rc rdb.Client, ctx context.Context, pattern string) ([]*T, error) {

	iter := rc.Scan(ctx, 0, pattern, 0).Iterator()

	var items []*T

	for iter.Next(ctx) {
		key := iter.Val()

		item, _ := Get[T](rc, ctx, key)

		items = append(items, item)
	}

	if err := iter.Err(); err != nil {
		return nil, iter.Err()
	}

	return items, nil

}

//func Update() {}

func Delete(rc rdb.Client, ctx context.Context, key string) error {
	_, err := rc.Del(ctx, key).Result()
	if err != nil {
		return err
	}

	return nil
}
