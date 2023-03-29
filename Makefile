GOPATH := $(shell go env GOPATH)
lint:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.0
	go install github.com/golangci/golangci-lint/cmd/golangci-lint
	$(GOPATH)/bin/golangci-lint run ./...
	go mod tidy
