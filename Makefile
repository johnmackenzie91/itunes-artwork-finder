start:
	docker-compose up -d

stop:
	docker-compose down
build:
	docker-compose build

logs:
	docker-compose logs -f
.PHONY: debug
debug:
	docker-compose -f ./debug/docker-compose.yml build
	docker-compose -f ./debug/docker-compose.yml up

test:
	go test -v -covermode=count -coverprofile=coverage.out ./...

integration-tests:
	go test ./tests/... -long -race

linter:
	golangci-lint run --concurrency 5 \
		--disable-all -E errcheck -E golint -E bodyclose -E gochecknoinits -E whitespace -E misspell \
		./...

gen-server:
	oapi-codegen -package="app" \
		-generate="types, chi-server" \
		--exclude-schemas=APIError \
		openapi.yaml  > ./internal/app/app.gen.go

coveralls:
	go tool cover -func=coverage.out
	${GOPATH}/bin/goveralls -coverprofile=coverage.out -repotoken ${COVERALLS_TOKEN}

gen-redoc:
	docker run -v $(shell pwd):/work \
		simplealpine/yaml2json:latest \
		/work/openapi.yaml > ./internal/app/redoc/docs/openapi.json
