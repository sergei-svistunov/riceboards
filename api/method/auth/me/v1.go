package me

import (
	"context"
	"strconv"

	"github.com/go-qbit/rbac"
	"github.com/go-qbit/rpc"

	"riceboards/authctx"
)

type reqV1 struct {
}

type userV1 struct {
	Id          uint32   `json:"id"`
	Fullname    string   `json:"fullname"`
	AvatarUrl   string   `json:"avatar_url"`
	Email       string   `json:"email"`
	Permissions []string `json:"permissions"`
}

var errorsV1 struct {
	Unauthorized rpc.ErrorFunc `desc:"Unauthorized"`
}

func (m *Method) ErrorsV1() interface{} {
	return &errorsV1
}

func (m *Method) V1(ctx context.Context, r *reqV1) (*userV1, error) {
	user := authctx.FromContext(ctx)
	if user == nil {
		return nil, errorsV1.Unauthorized("Invalid user")
	}

	u, err := m.db.Users.GetUserById(ctx, strconv.FormatUint(uint64(user.Id), 10))
	if err != nil {
		return nil, err
	}

	roles, err := rbac.GetUserRolesIds(ctx, nil)
	if err != nil {
		return nil, err
	}

	rolesPermissions, err := rbac.GetRolesPermissions(ctx, roles...)
	if err != nil {
		return nil, err
	}

	return &userV1{
		Id:          u.(authctx.User).Id,
		Fullname:    u.(authctx.User).Fullname,
		Email:       u.(authctx.User).Email,
		AvatarUrl:   u.(authctx.User).AvatarUrl,
		Permissions: rolesPermissions,
	}, nil
}
