package api

import (
	"context"

	"github.com/godverv/hello_world/pkg/api"
)

func (a *Api) Get(ctx context.Context, r *api.Get_Request) (*api.Value, error) {

	return &api.Value{}, nil
}
