yaml_file = ./files/yml/apiTest.local.yml
cmd_folder = ./cmd/neo_post/
gorun = @go run

run:
	${gorun} ${cmd_folder}cmd -config ${yaml_file}

migration:
	${gorun} ${cmd_folder}migration -config ${yaml_file}