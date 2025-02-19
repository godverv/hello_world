package api

import (
	"context"

	errors "github.com/Red-Sock/trace-errors"

	api "github.com/godverv/hello_world/pkg/hello_world"
)

func (a *Impl) Set(ctx context.Context, r *api.Set_Request) (*api.Set_Response, error) {
	prep, err := a.db.PrepareContext(ctx, `
		INSERT INTO user_values 
			(key, value)
		VALUES ($1, $2)
		ON CONFLICT (key) DO UPDATE SET value = EXCLUDED.value
`)
	if err != nil {
		return nil, errors.Wrap(err, "error creating prepare")
	}

	defer prep.Close()

	for _, kv := range r.Vals.Values {
		_, err = prep.ExecContext(ctx, kv.Key, kv.Value)
		if err != nil {
			return nil, errors.Wrap(err, "error inserting key value")
		}
	}

	return &api.Set_Response{}, nil
}
