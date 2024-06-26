GO_VERSION := 1.20
.PHONY: install-go init-go

setup: install-go init-go copy-hooks install-lint

install-go:
	brew install "go@$(GO_VERSION)"

init-go:
	echo 'export PATH=$$PATH:/usr/local/go/bin/go' >> $${HOME}/.zshrc
	echo 'export PATH=$$PATH:$${HOME}/go/bin' >> $${HOME}/.zshrc

install-lint:
	brew install golangci-lint

copy-hooks:
	chmod +x scripts/hooks/*
	cp -r scripts/hooks .git/.

build:
	go build -o api cmd/main.go

test:
	go test ./... -coverprofile=coverage.out

coverage:
	go tool cover -func coverage.out | grep "total:" | awk '{print ((int($$3) > 80) != 1) }'

report:
	go tool cover -html=coverage.out -o cover.html

check-format:
	test -z $$(go fmt ./...)

static-check:
	golangci-lint run