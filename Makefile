yaml_file = ./files/yml/apiTest.local.yml
cmd_folder = ./cmd/neo_post/
gorun = @go run
gobuild = @go build

run:
	${gorun} ${cmd_folder}cmd -config ${yaml_file}

migration:
	${gorun} ${cmd_folder}migration -config ${yaml_file}

drop:
	${gorun} ${cmd_folder}drop -config ${yaml_file}

restart: drop migration run

cli:
	${gorun} ${cmd_folder}ruler

run-docker:
	docker compose up --build -d

stop-docker:
	docker compose down

run-docker-migration:
	docker compose up --build -d neo-post-migration

BIN=./bin

MAIN_BIN=main
RULER_BIN=ruler
DROPER_BIN=droper
MIGRATOR_BIN=migrator

build-main:
	${gobuild} -o $(BIN)/$(MAIN_BIN) ./cmd/neo_post/cmd/*.go

run-main:
	./$(BIN)/$(MAIN_BIN)

clean-main:
	rm -f ./$(BIN)/$(MAIN_BIN)

refresh-main: clean-main build-main run-main

build-droper:
	${gobuild} -o $(BIN)/$(DROPER_BIN) ./cmd/neo-post/drop/

clean-droper:
	rm -f ./$(BIN)/$(DROPER_BIN)

run-droper:
	./$(BIN)/$(DROPER_BIN)

refresh-droper: clean-droper build-droper run-droper

build-migrator:
	${gobuild} -o $(BIN)/$(MIGRATOR_BIN) ./cmd/neo-post/migration/

run-migrator:
	./$(BIN)/$(MIGRATOR_BIN)

clean-migrator:
	rm -f ./$(BIN)/$(MIGRATOR_BIN)

refresh-migrator: clean-migrator build-migrator run-migrator

build-ruler:
	${gobuild} -o $(BIN)/$(RULER_BIN) ./cmd/neo_post/ruler/*.go

run-ruler: build-ruler
	./$(BIN)/$(RULER_BIN) $(ARGS)

clean-ruler:
	rm -f ./$(BIN)/(RULER_BIN)

refresh-ruler: clean-ruler build-ruler run-ruler

