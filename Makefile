.PHONY: test

install:
	go install ./...

lint:
ifndef STATICCHECK
	go get -u honnef.co/go/tools/cmd/staticcheck
endif
	go vet ./...
	staticcheck ./...

test: lint
	go test -race ./... -timeout 1s

release: install test
	bump_version minor main.go
