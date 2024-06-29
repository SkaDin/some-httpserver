export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgresql://test:test@localhost:5432/testDB?sslmode=disable


Запуск гуся:
goose -dir db/migrations up