GOCMD = go
GOSWAG = swag
GOSWAGGER = swagger
MIGRATE = migrate
MIGRATION_PATH = ecommerce.db
DATABASE_URL = "sqlite3://ecommerce.db?sslmode=disable&x-migrations-table=ecommerce_schema_migrations"

check-air:
	which ${AIRCMD} || (echo "please install go-swagger from https://github.com/cosmtrek/air" && exit 1)

serve: check-air
	${AIRCMD}

run:
	${GOCMD} run main.go

check-goswag:
	which ${GOSWAG} || (${GOCMD} install github.com/swaggo/swag/cmd/swag@v1.7.8)

check-goswagger:
	which ${GOSWAGGER} || (echo "please install go-swagger from https://goswagger.io/install.html" && exit 1)

swagger: check-goswag
	GO111MODULE=on ${GOSWAG} init -g cmd/router.go -o doc

swagger-md: check-goswagger
	GO111MODULE=on ${GOSWAG} init -g cmd/router.go -o doc
	GO111MODULE=on ${GOSWAGGER} generate markdown -f doc/swagger.json --output doc/api.md

#MIGRATION RELATED
check-migrate:
	which ${MIGRATE} || brew install golang-migrate

#migrate-up command used to run migration up
migrate-up: check-migrate
	@echo "Running Migration"
	${MIGRATE} -source file://${MIGRATION_PATH} -database "$DATABASE_URL" up

#new-migration command used to create new migration file in migrations/sql dir
#example make new-migration name="filename".
#it will generate 2 sql file for migration up and down
new-migration: check-migrate
	${MIGRATE} create -ext sql -dir ${MIGRATION_PATH} -seq $(name)
