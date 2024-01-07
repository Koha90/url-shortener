GO_VERSION := 1.21

.PHONY: install-go init-go

setup: install-go init-go

# TODO add MacOS support
install-go:
	wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz"
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
	rm go$(GO_VERSION).linux-amd64.tar.gz

init-go:
	echo 'export PATH=$$PATH:/usr/local/go/bin' >> $${HOME}/.bashrc
	echo 'export PATH=$$PATH:${HOME}/go/bin' >> $${HOME}/.bashrc

build:
	@go build -o bin/url-shortener cmd/main.go

run: build
	@./bin/url-shortener

test:
	@go test ./... -coverprofile=coverage.out

coverage:
	@go tool cover -func coverage.out | grep "total:" | \
		awk '{print ((int($$3) > 80) != 1) }'

report: 
	@go tool cover -html=coverage.out -o cover.xhtml
