PORT=8080
DSN="host=localhost port=5432 user=root password=admin dbname=chigodb sslmode=disable timezone=UTC connect_timeout=5"

DB_DOCKER_CONTAINER=chigo
BINARY_NAME=chi-go
DB_POSTGRES_NAME=chigodb

# Untuk menjalakan makefile dapat dilakukan dengan make {method dibawah}. contoh make postgres
# Makefile hanya berisikan perintah untuk menjalakan command di cmd ataupun terminal

postgres:
	@docker run --name ${DB_DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_PASSWORD=admin -e POSTGRES_USER=root -d postgres

createdb:
	@docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=root --owner=root chigodb

start_docker:
	@docker start ${DB_DOCKER_CONTAINER}

stop_containers:
	@echo "Stopping Container"
	@if [ "$(shell docker ps -q)" ]; then \
		echo "Found and stopped Containers"; \
		docker stop $(shell docker ps -q); \
	else \
		echo "No Active containers found.."; \
	fi

# menggukana sqlx untuk membuat migrasi. sqlx migrate add -r (nama migrasi) yang berisikan pembuatan table dll (query)
create_migrations:
	@sqlx migrate add -r init 

# terdapat sqlx untuk migrasi up dan down yang bisa dilihat pada folder migrasi
migrate-up:
	@sqlx migrate run --database-url "postgres://root:admin@localhost:5432/chigodb?sslmode=disable"

migrate-down:
	@sqlx migrate revert --database-url "postgres://root:admin@localhost:5432/chigodb?sslmode=disable"

build:
	@echo "Building API Binary"
	@go build -o ${BINARY_NAME} cmd/*.go
	@echo "Binary Built Successfully"

run: build stop_containers start_docker
	@echo "Starting API"
	@env PORT=${PORT} DSN=${DSN} ./${BINARY_NAME} &
	@echo "API started!"


stop:
	@echo "Stopping backend"
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped backend"

start: run 

restart : stop start
