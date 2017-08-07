create-pre-commit-hooks: ## creates pre-commit hooks
	chmod +x $(CURDIR)/hooks/pre-commit
	ln -s $(CURDIR)/hooks/pre-commit .git/hooks/pre-commit || true

deps:
	glide install

.PHONY: deps
