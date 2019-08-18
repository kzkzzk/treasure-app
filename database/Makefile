export GO111MODULE := off

<<<<<<< HEAD
DBNAME:=j-sweet
=======
DBNAME:=treasure_app
>>>>>>> master
DATASOURCE:=root:password@tcp(127.0.0.1:3306)/$(DBNAME)

install:
	which goose || GO111MODULE=off go get -u github.com/pressly/goose/cmd/goose

docker-compose/up:
	docker-compose up -d

docker-compose/down:
	docker-compose down

mysql:
	mysql -u root -h localhost --protocol tcp -p $(DBNAME)

migrate/init:
	mysql -u root -h localhost --protocol tcp -e "create database \`$(DBNAME)\`" -p

migrate/status:
	@goose -dir migrations mysql "$(DATASOURCE)" status

<<<<<<< HEAD
migrate/new:
	@goose -dir migrations mysql "$(DATASOURCE)" create "$(NAME)" sql

=======
>>>>>>> master
migrate/up:
	@goose -dir migrations mysql "$(DATASOURCE)" up

__migrate/down:
	@goose -dir migrations mysql "$(DATASOURCE)" down

<<<<<<< HEAD
init: install docker-compose/up migrate/init migrate/up
=======
init: install docker-compose/up migrate/init migrate/up
>>>>>>> master
