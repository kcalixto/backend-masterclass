#--------- START DOCKER ---------#

# 1st
create-docker:
	docker run --name postgresdb -p 5432:5432 -e POSTGRES_PASSWORD=123 -e POSTGRES_USER=root -d postgres:12-alpine


#--------- CREATE MIGRATIONS FILES ---------#

# 2nd
create-migrations:
	migrate create -ext sql -dir db/migration -seq init_schema


#--------- CREATE TABLE ---------#

# 3rd
exec-docker:
	docker exec -it postgresdb /bin/sh
# 3rd.1 (inside docker container): createdb --username=root --owner=root simple_bank
# 3rd.2 (inside docker container): psql simple_bank

# OR

#3rd
exec-docker-2:
	docker exec -it postgresdb createdb --username=root --owner=root simple_bank


#--------- INSTALL SQLC ---------#
install-sqlc:
	brew install sqlc


mock:
	mockgen -package mockdb -destination db/mock/store.go backend-masterclass/db/sqlc Store