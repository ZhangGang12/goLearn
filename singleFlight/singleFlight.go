package main

import (
	"context"
	"errors"

	"golang.org/x/sync/singleflight"
)

var SingleF singleflight.Group

func LoadMaterial(ctx context.Context, k string, fn func() (interface{}, error)) (interface{}, error) {
	resultChan := SingleF.DoChan(k, fn)
	select {
	case c := <-resultChan:
		return c.Val, c.Err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ctx.Done()
	cancel()
	_, _ = LoadMaterial(ctx, "test", func() (interface{}, error) {
		return nil, errors.New("123123")
	})

	_, _ = LoadMaterial(ctx, "test1", func() (interface{}, error) {
		return nil, nil
	})

}
