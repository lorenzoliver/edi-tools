lint:
	golangci-lint run


fmt:
	golangci-lint fmt

test:
	go test ./... -coverprofile=cover.out
