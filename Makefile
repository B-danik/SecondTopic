createContainer:
	docker run --name-todo-name -e POSTGRES_PASSWORD='D112233Diss' -p 5432:5432 -d --rm postgres
createMigrate:
	migrate create -ext sql -dir .\schema -seq init
createDB:
	docker exec -it caf4ccb71783 /bin/bash
migrateUp:
	migrate -path ./schema -database 'postgres://postgres:D112233Diss@localhost:5432/postgres?sslmode=disable' up
migrateDown:
	migrate -path ./schema -database 'postgres://postgres:D112233Diss@localhost:5432/postgres?sslmode=disable' down	