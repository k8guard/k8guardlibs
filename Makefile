.DEFAULT_GOAL := help
.PHONY: help

create-pre-commit-hooks: ## creates pre-commit hooks
	chmod +x $(CURDIR)/hooks/pre-commit
	ln -s $(CURDIR)/hooks/pre-commit .git/hooks/pre-commit || true

deps: ## install dependencies
	glide install

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
