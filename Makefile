# Makefile at the repo root
.PHONY: mock

## Generate all mocks from the root .mockery.yaml
mock:
	mockery
