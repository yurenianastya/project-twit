install-lint:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.39.0

run-linter:
	sudo $(shell go env GOPATH)/bin/golangci-lint run

test:
	cd $(shell go env GOPATH)/src/project-twit/methods && go test -cover
