package docs

import (
	_ "embed"
	"net/http"

	swaggerui "github.com/alexliesenfeld/go-swagger-ui"
)

//go:embed hello_world_api.swagger.json
var spec []byte

const basePath = "/docs/"

func Swagger() (path string, handler http.HandlerFunc) {
	return basePath, swaggerui.NewHandler(
		swaggerui.WithBasePath(basePath),
		swaggerui.WithHTMLTitle("HELLO_WORLD_DOCS"),
		swaggerui.WithSpec(spec),
		swaggerui.WithTryItOutEnabled(true),
		swaggerui.WithPersistAuthorization(true),
	)
}
