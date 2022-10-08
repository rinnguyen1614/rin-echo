APP_PREFIX=rin
DB_URL=postgresql://root:secret@localhost:15432/rin-echo?sslmode=disable
DB_PORT=15432
DB_NAME=rin-echo
DB_USER=root
DB_PASSWORD=secret
DIR_MIGRATE=./internal/system/db/migrate

network:
	docker network create ${APP_PREFIX}-network

postgres:
	docker run --name ${APP_PREFIX}-postgres --network ${APP_PREFIX}-network -p ${DB_PORT}:5432 -e POSTGRES_USER=${DB_USER} -e POSTGRES_PASSWORD=${DB_PASSWORD}

createdb:
	docker exec -it ${APP_PREFIX}-postgres createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}

dropdb:
	docker exec -it ${APP_PREFIX}-postgres dropdb ${DB_NAME}

migrate-up:
	migrate -path ${DIR_MIGRATE} -database ${DB_URL} -verbose up

migrate-down:
	migrate -path ${DIR_MIGRATE} -database ${DB_URL} -verbose down

server-generate:
	cd ./internal/system && go generate

server:
	cd ./internal/system && go run main.go

web:
	cd ./web && npm start