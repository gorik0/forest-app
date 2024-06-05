
POSTGRES_USER = postgres
MIGRATE_NAME = postgres
POSTGRES_PASSWORD = 1
POSTGRES_HOST = localhost
POSTGRES_PORT = 5432
POSTGRES_DB = postgres


create_migration:
	migrate create -ext=sql -dir="/Users/gorik/go/src/awesomeProject/migrations" -seq ${MIGRATE_NAME}

migrate_up:
	migrate -path="/Users/gorik/go/src/awesomeProject/migrations" -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up

migrate_down:
	migrate -path="/Users/gorik/go/src/awesomeProject/migrations" -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose down


migrate_fix:
	migrate -path="/Users/gorik/go/src/awesomeProject/migrations" -database="postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" force 1


gre-docker:
	docker run -it --rm  -p 5432:5432 \
 	--name pgsql-dev \
 	-e POSTGRES_PASSWORD=1 \
 	-v /data:/var/lib/postgresql/data \
   	postgres

docker-clean: d-stop d-c-cont d-clean_image
d-stop:
	docker stop $(shell docker ps -aq)

d-c-cont:
	docker rm $(shell docker ps -aq)
d-clean_image:
	docker rmi $(shell docker image ls -aq)

.PHONY: create_migration migrate_up migrate_down