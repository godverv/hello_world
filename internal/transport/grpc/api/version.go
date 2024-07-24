package api

import (
	"context"

	"github.com/godverv/hello_world/pkg/api"
)

func (a *Api) Version(_ context.Context, _ *api.Version_Request) (
	*api.Version_Response, error) {

	return &api.Version_Response{
		Version: a.version,
	}, nil
}
