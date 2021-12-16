package authctx

import (
	"context"
	"fmt"
)

type ctxKeyT struct{}

type User struct {
	Id        uint32
	Fullname  string
	AvatarUrl string
	Email     string
	IsAdmin   bool
}

var (
	ctxKey ctxKeyT

	ErrUnauthorized = fmt.Errorf("unathorized")
)

func ToContext(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, ctxKey, user)
}

func FromContext(ctx context.Context) *User {
	v := ctx.Value(ctxKey)
	if v == nil {
		return nil
	}

	return v.(*User)
}

func GetCurUserId(ctx context.Context) (uint32, error) {
	curUser := FromContext(ctx)
	if curUser == nil {
		return 0, ErrUnauthorized
	}

	return curUser.Id, nil
}
