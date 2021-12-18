package google

import (
	"context"
	"log"
	"strconv"

	"github.com/go-qbit/rbac"
	"github.com/go-qbit/rpc"
	"google.golang.org/api/idtoken"

	"riceboards/authctx"
	"riceboards/db/users"
)

type reqV1 struct {
	Credential string `json:"credential"`
}

type userV1 struct {
	Id          uint32   `json:"id"`
	Email       string   `json:"email"`
	Fullname    string   `json:"fullname"`
	AvatarUrl   string   `json:"avatar_url"`
	Token       string   `json:"token"`
	Permissions []string `json:"permissions"`
}

var errorsV1 struct {
	InvalidToken rpc.ErrorFunc `desc:"Invalid token"`
}

func (m *Method) ErrorsV1() interface{} {
	return &errorsV1
}

func (m *Method) V1(ctx context.Context, r *reqV1) (*userV1, error) {
	p, err := idtoken.Validate(ctx, r.Credential, "")
	if err != nil {
		return nil, errorsV1.InvalidToken(err.Error())
	}

	log.Printf("%#v", p.Claims)

	uid, err := m.db.Users.Register(ctx, &users.OAuthUser{
		Fullname:  p.Claims["name"].(string),
		Email:     p.Claims["email"].(string),
		AvatarUrl: p.Claims["picture"].(string),
	})

	user, err := m.db.Users.GetUserById(ctx, strconv.FormatUint(uint64(uid), 10))
	if err != nil {
		return nil, err
	}
	u := user.(authctx.User)
	ctx = authctx.ToContext(ctx, &u)

	roles, err := rbac.GetUserRolesIds(ctx, nil)
	if err != nil {
		return nil, err
	}

	rolesPermissions, err := rbac.GetRolesPermissions(ctx, roles...)
	if err != nil {
		return nil, err
	}

	return &userV1{
		Id:          uid,
		Fullname:    p.Claims["name"].(string),
		Email:       p.Claims["email"].(string),
		AvatarUrl:   p.Claims["picture"].(string),
		Token:       m.auth.GetSessionId(uid),
		Permissions: rolesPermissions,
	}, nil
}
