up:
	docker-compose -f deployments/docker-compose.yaml up -d

test:
	go test -race ./internal/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.33.0

lint: install-lint-deps
	golangci-lint run ./...

migrate:
	docker-compose -f deployments/docker-compose.yaml exec app sh -c "/opt/social/migrations --config=/etc/social/config.yaml"

test-data:
	docker-compose -f deployments/docker-compose.yaml exec app sh -c "/opt/social/testdatagen --config=/etc/social/config.yaml"

swagger-init:
	go get -u github.com/swaggo/swag/cmd/swag

swagger: swagger-init
	swag init -g ./internal/server/http/router.go -o api