citest:
	go test ./... -v -coverprofile c.out
	export GIT_COMMITED_AT=$(date + "%s")

test:
	mkdir -p coverage
	go test ./... -v -coverprofile ./coverage/.coverage
	go tool cover -html=./coverage/.coverage -o ./coverage/index.html
