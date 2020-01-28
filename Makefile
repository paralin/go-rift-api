SHELL=/bin/bash
SWAGGER=hack/bin/swagger
GOLIST=go list -f "{{ .Dir }}" -m
GO111MODULE=on

all:

vendor:
	go mod vendor

$(SWAGGER):
	cd ./hack; \
	go build -v \
		-o ./bin/swagger \
		github.com/go-swagger/go-swagger/cmd/swagger

gengo: $(SWAGGER)
	shopt -s globstar; \
	set -eo pipefail; \
	export PROJECT=$$(go list -m); \
	export PATH=$$(pwd)/hack/bin:$${PATH}; \
	mkdir -p ./client; \
	$(SWAGGER) generate client \
		--skip-validation \
		-f ./resources/rift-api/swagger.json \
		-A lol \
		--tags=Plugins \
		-t ./

validate: $(SWAGGER)
	$(SWAGGER) validate ./resources/rift-api/swagger.json

test:
	go test -v ./...


