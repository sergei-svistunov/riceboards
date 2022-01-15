package projects

import (
	"context"

	"github.com/go-qbit/model"
	"github.com/go-qbit/model/expr"
)

func (t *Table) GetIdByStrId(ctx context.Context, strId string) (uint32, error) {
	data, err := t.GetAll(ctx, []string{"id"}, model.GetAllOptions{
		Filter: expr.Eq(t.FieldExpr("str_id"), expr.Value(strId)),
		Limit:  1,
	})
	if err != nil {
		return 0, err
	}

	if data.Len() == 0 {
		return 0, nil
	}

	return data.Data()[0][0].(uint32), nil
}
