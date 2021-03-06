# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

GOLANGCI_VERSION = 1.9.3

.PHONY: clean
clean: ## Clean the working area
	rm -rf bin/ build/ vendor/

.PHONY: check
check: test lint ## Run tests and linters

.PHONY: test
test: ## Run tests
	go test ${ARGS} ./...

bin/golangci-lint: ## Install golangci linter
	@mkdir -p ./bin/
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- v${GOLANGCI_VERSION}

.PHONY: lint
lint: bin/golangci-lint ## Run linter
	@bin/golangci-lint run

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Variable outputting/exporting rules
var-%: ; @echo $($*)
varexport-%: ; @echo $*=$($*)
