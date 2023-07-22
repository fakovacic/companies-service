.PHONY: lint
lint:
	@golangci-lint run

.PHONY: test
test:
	@go test -race -cover ./internal/...

.PHONY: compile
compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o bin/companies ./cmd/companies/main.go

.PHONY: docker-build
docker-build:
	docker build -t companies .

.PHONY: compose-start
compose-start:
	@docker-compose up -d

.PHONY: compose-stop
compose-stop:
	@docker-compose down

.PHONY: compose-remove
compose-remove:
	@docker-compose rm -s -f

