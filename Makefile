#================================
#== DOCKER ENVIRONMENT
#================================


include .env

COMPOSE := @docker-compose

dcb:
	${COMPOSE} build

dcuf:
ifdef f
	${COMPOSE} up -d --${f}
endif

dcubf:
ifdef f
	${COMPOSE} up -d --build --${f}
endif

dcu:
	${COMPOSE} up -d --build

dcd:
	${COMPOSE} down

#================================
#== GOLANG ENVIRONMENT
#================================
GO := @go
GIN := @gin

goinstall:
	${GO} get .

godev:
	${GIN} -a 4000 -p 3001 -b bin/main run main.go

# goprod:
# 	${GO} build -o main .

gotest:
	${GO} test -v

goformat:
	${GO} fmt ./...

create_migration:
	migrate create -ext=sql -dir=db/migration -seq init 

migrate_up:
		migrate -path=db/migration -database "${DATABASE_URI_DEV}?sslmode=disable" -verbose up

migrate_down:
		migrate -path=db/migration -database "${DATABASE_URI_DEV} -verbose down

.PHONY: create_migration migrate_up migrate_down