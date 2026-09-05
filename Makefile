include .env
export
export PROJECT_ROOT=$(shell pwd)

run:
	go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/todoapp/main.go

docker-up:
	docker compose up -d

docker-down:
	docker compose down

clean-up:
	@read -p "Очистить все volume файлы окружений? Данные будут утеряны. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down todoapp-postgres port-forwarder && \
		rm -rf out/pgdata && \
		echo "Файлы очищены"; \
	else \
		echo "Очистка отменена"; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутствует параметр seq. make migrate-create seq=название-миграции"; \
		exit 1; \
	fi
	docker compose run --rm todoapp-postgres-migrate create -ext sql -dir /migrations -seq "$(seq)"
migrate-up:
	@if [ -z "$(level)" ]; then \
		docker compose run --rm todoapp-postgres-migrate \
			-path /migrations \
			-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DATABASE}?sslmode=disable \
			up; \
	else \
		docker compose run --rm todoapp-postgres-migrate \
			-path /migrations \
			-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DATABASE}?sslmode=disable \
			up "$(level)"; \
	fi

migrate-down:
	@if [ -z "$(level)" ]; then \
		docker compose run --rm todoapp-postgres-migrate \
			-path /migrations \
			-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DATABASE}?sslmode=disable \
			down; \
	else \
		docker compose run --rm todoapp-postgres-migrate \
			-path /migrations \
			-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DATABASE}?sslmode=disable \
			down "$(level)"; \
	fi

todoapp-deploy:
	docker compose up -d --build todoapp

ps:
	docker compose ps

swagger-gen:
	@docker compose run --rm swagger \
		init \
		-g cmd/todoapp/main.go \
		-o docs \
		--parseInternal \
		--parseDependency