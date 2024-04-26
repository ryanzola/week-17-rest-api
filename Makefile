build:
	@go build -o bin/api

run: build
	@./bin/api

seed:
	@go run scripts/seed.go

docker:
	echo "Building Dockerfile..."
	@docker build -t api .
	echo "Running API inside Docker container..."
	@docker run -p 3001:3001 api

test:
	@go test -v ./...