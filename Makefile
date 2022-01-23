.PHONY: citest
citest:
	ENV=test go test ./... -v -coverprofile c.out
	export GIT_COMMITED_AT=$(date + "%s")

.PHONY: test
test:
	ENV=test echo $(ENV)
	mkdir -p coverage
	ENV=test go test ./... -v -coverprofile ./coverage/.coverage
	go tool cover -html=./coverage/.coverage -o ./coverage/index.html
