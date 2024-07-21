yaml_file = ./files/yml/apiTest.local.yml
cmd_folder = ./cmd/neo_post/
gorun = @go run

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