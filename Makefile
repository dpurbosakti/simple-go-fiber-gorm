## createdb: create database
createdb:
	docker exec -it postgres14 createdb --username=root --owner=root fiber_gorm
## postgresrun: start docker postgres
postgresrun:
	docker start postgres14

## test: run test
test:
	go test -v -cover ./...

## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: createdb postgresrun test help 