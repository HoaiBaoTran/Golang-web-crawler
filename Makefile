DATABASE_URL ?= "postgresql://postgres:hoaibao1142@localhost:5432/web_crawler?sslmode=disable"

MIGRATION_PATH ?= migrations

migrate-create:
	migrate create -ext=sql -dir=${MIGRATION_PATH} ${name}

migrate-v1-force:
	migrate -path=${MIGRATION_PATH} -database=${DATABASE_URL} force 1

migrate-up-v1:
	migrate -path=${MIGRATION_PATH} -database=${DATABASE_URL} -verbose up 1

migrate-down-v1:
	migrate -path=${MIGRATION_PATH} -database=${DATABASE_URL} -verbose down 1