package users

import (
	"context"

	"github.com/go-qbit/model"
	"github.com/go-qbit/model/expr"
	"github.com/go-qbit/qerror"

	"riceboards/authctx"
)

type OAuthUser struct {
	Fullname  string
	Email     string
	AvatarUrl string
}

func (t *Table) Register(ctx context.Context, user *OAuthUser) (uint32, error) {
	var uid uint32

	return uid, t.GetDb().DoInTransaction(ctx, func(ctx context.Context) error {
		var users []struct {
			Id uint32
		}
		if err := t.GetAllToStruct(ctx, &users, model.GetAllOptions{
			Filter:    expr.Eq(t.FieldExpr("email"), expr.Value(user.Email)),
			Limit:     1,
			ForUpdate: true,
		}); err != nil {
			return err
		}

		if len(users) == 1 {
			uid = users[0].Id
			return nil
		}

		pks, err := t.AddMulti(ctx, model.NewData(
			[]string{"fullname", "email", "avatar_url", "is_admin"},
			[][]interface{}{
				{user.Fullname, user.Email, user.AvatarUrl, false},
			},
		), model.AddOptions{})

		if err != nil {
			return err
		}

		uid = pks.Data()[0][0].(uint32)

		return nil
	})
}

func (t *Table) GetUserById(ctx context.Context, id string) (interface{}, error) {
	var res []authctx.User
	if err := t.GetAllToStruct(ctx, &res, model.GetAllOptions{
		Filter: expr.Eq(t.FieldExpr("id"), expr.Value(id)),
		Limit:  1,
	}); err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, qerror.Errorf("User does not exist")
	}

	return res[0], nil
}
