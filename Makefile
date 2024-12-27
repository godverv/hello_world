include rscli.mk


dep:
	easyp mod download


easyp:
	mkdir -p pkg/web
	mkdir -p pkg/docs/
	easyp generate