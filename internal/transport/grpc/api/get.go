package api

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	api "github.com/godverv/hello_world/pkg/hello_world"
)

func (a *Impl) Get(ctx context.Context, r *api.Get_Request) (*api.Value, error) {
	res := &api.Value{
		Key: r.Key,
	}

	err := a.db.QueryRowContext(ctx, `
			SELECT 
				value 
			FROM user_values
			WHERE key = $1
`, r.Key).
		Scan(&res.Value)
	if err != nil {
		return nil, errors.Wrap(err, "error reading from db")
	}

	return res, nil
}
