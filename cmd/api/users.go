package main

import (
	"context"

	"github.com/iangechuki/go-ecommerce/internal/store"
)

type userKey string

var userCtx userKey

func getUserFromContext(ctx context.Context) *store.User {
	return ctx.Value(userCtx).(*store.User)
}
