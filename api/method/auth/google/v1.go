package google

import (
	"context"
	"log"

	"github.com/go-qbit/rpc"
	"google.golang.org/api/idtoken"

	"riceboards/db/users"
)

type reqV1 struct {
	Credential string `json:"credential"`
}

type userV1 struct {
	Id        uint32 `json:"id"`
	Email     string `json:"email"`
	Fullname  string `json:"fullname"`
	AvatarUrl string `json:"avatar_url"`
	Token     string `json:"token"`
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

	return &userV1{
		Id:        uid,
		Fullname:  p.Claims["name"].(string),
		Email:     p.Claims["email"].(string),
		AvatarUrl: p.Claims["picture"].(string),
		Token:     m.auth.GetSessionId(uid),
	}, nil
}
