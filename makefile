tests:
	go test -race -timeout 30s -cover ./cmd ./check
	GO111MODULE=on go test -v -short -race -timeout 30s -coverprofile=coverage.txt -covermode=atomic ./...