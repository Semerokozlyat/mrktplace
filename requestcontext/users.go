package requestcontext

import (
	"context"

	"mrktplace/models"
)

type ctxKey int

const (
	userKey ctxKey = 1
)

func WithUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func User(ctx context.Context) *models.User {
	val := ctx.Value(userKey)
	u, ok := val.(*models.User)
	if !ok {
		return nil
	}
	return u
}
