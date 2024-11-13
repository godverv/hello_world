gen: gen-server-grpc

build-local-container:
	docker buildx build \
			--load \
			--platform linux/arm64 \
			-t hello_world:local .


### Grpc server generation
gen-server-grpc: .pre-gen-server-grpc .deps-grpc .gen-server-grpc

.pre-gen-server-grpc:
	mkdir -p pkg/
	mkdir -p pkg/docs
	mkdir -p pkg/web

.deps-grpc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/Red-Sock/protoc-gen-docs@v0.0.1

	rm -rf api/google
	rm -rf api/validate
	ln -sf $(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis/google api/google
	ln -sf $(GOPATH)/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v1.0.2/validate api/validate

.gen-server-grpc:
	protoc \
    	-I=./api \
    	-I $(GOPATH)/bin \
    	--openapiv2_out=./pkg/docs/swaggers \
    	--docs_out=./pkg/docs \
    	--go-grpc_out=./pkg/ \
    	--grpc-gateway_out=logtostderr=true:./pkg/ \
    	--grpc-gateway-ts_out=./pkg/web \
    	--go_out=./pkg/ \
	    --validate_out="lang=go:./pkg" \
    	./api/grpc/*.proto