TEST_FLAGS = -v -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.txt
GOLANGCI_LINT = bin/golangci-lint

.PHONY: test
test:
	go test $(TEST_FLAGS) ./...

.PHONY: setup
setup:
	go mod download

$(GOLANGCI_LINT):
	curl -sSL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh

.PHONY: lint
lint: $(GOLANGCI_LINT)
	./$(GOLANGCI_LINT) run ./...

.PHONY: cover
cover: test
	go tool cover -html=coverage.txt
