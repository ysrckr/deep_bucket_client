build:
	@cd cmd/web && bun build
	@go build -buildvcs=false -o ./bin/deepbucketclient ./cmd/api/main.go
	@cp .env.production ./bin/.env

dev:
	@cd cmd/web && bun dev & air && fg

clean:
	@rm -rf ./bin
	@rm -rf ./tmp
	@lsof -t -i:8000 | xargs kill
	@lsof -t -i:5173 | xargs kill
	@lsof -t -i:5174 | xargs kill

.PHONY: build dev clean