dockerPull:
	docker pull postgres
createContainer:
	docker run --name=todo-name -e POSTGRES_PASSWORD='D112233Diss' -p 5432:5432 -d --rm postgres
createMigrate:
	migrate create -ext sql -dir .\schema -seq init
execDB:
	docker exec -it 3d9a1631e672 /bin/bash
migrateUp:
	migrate -path ./internal/database/postgre/schema -database 'postgres://postgres:D112233Diss@localhost:5432/postgres?sslmode=disable' up
migrateDown:
	migrate -path ./internal/database/postgre/schema -database 'postgres://postgres:D112233Diss@localhost:5432/postgres?sslmode=disable' down