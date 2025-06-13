BINARY_NAME=test-task
DB_URL=postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable

build:
	go build -o $(BINARY_NAME) ./cmd/main.go

run:
	go run ./cmd/main.go

migrate-up:
	migrate -path ./migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path ./migrations -database "$(DB_URL)" down

migrate-create:
	migrate create -ext sql -dir ./migrations -seq new_migration

docker-build:
	docker-compose build

docker-up:
	docker-compose up

docker-down:
	docker-compose down

clean:
	rm -f $(BINARY_NAME)