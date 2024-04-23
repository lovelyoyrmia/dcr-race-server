DB_URL=mysql://root:root@tcp(localhost:3306)/dcr?parseTime=true

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_PASSWORD=root -e POSTGRES_USER=root -d postgres:12-alpine

server:
	go run cmd/main.go

dockerbuild:
	docker build -t dcr:latest .

seed:
	go run seeds/seed.go

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root dcr

dropdb:
	docker exec -it postgres12 dropdb dcr

migrateup:
	migrate -path migrations -database "$(DB_URL)" --verbose up

createmigrate:
	migrate create -ext sql -dir migrations -seq init_schema

migratedown:
	migrate -path migrations -database "$(DB_URL)" --verbose down

newmigration:
	migrate create -ext sql -dir migrations -seq $(name)

sqlc:
	sqlc generate
