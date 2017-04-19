.PHONY: test

STATICCHECK := $(shell command -v staticcheck)

install:
	go install ./...

test:
	go test -race ./... -timeout 1s

release: install test
	bump_version minor main.go
