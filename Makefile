#@ Test
lint: ## Run linters only.
	@echo -e "\033[2m→ Running linters...\033[0m"
	golangci-lint run --config .golangci.yml

test: ## Run go tests for files with tests.
	@echo -e "\033[2m→ Run tests for all files...\033[0m"
	go test --cover ./...

check: lint test ## Run full check: lint and test.

##@ Generate
mocks: ## Generate mocks.
	@echo -e "\033[2m→ Generating mocks...\033[0m"
	mockgen -source=client.go -destination=mocks/client_mock.go -package=mocks
	
##@ Other
#------------------------------------------------------------------------------
help:  ## Display help
	@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
#------------- <https://suva.sh/posts/well-documented-makefiles> --------------

.DEFAULT_GOAL := help
.PHONY: help lint test check mocks
