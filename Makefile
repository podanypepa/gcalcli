.DEFAULT_GOAL := lint
APPNAME=$(shell basename "$(PWD)")
REPO=$(shell go list)
BUILDINFO="./pkg/version/buildinfo.txt"
BIN_LINUX="./$(APPNAME)-linux"
BIN_WINDOWS="./$(APPNAME).exe"
BIN_DARWIN_ARM="./$(APPNAME)-darwin.arm64"
BIN_DARWIN_X8664="./$(APPNAME)-darwin.x86_64"

.PHONY: help

all: help

help: Makefile
	@echo
	@echo " Choose a command run in:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## lint: checke all sources for errors
lint: getversion
	wrapcheck ./...
	revive ./...
	errcheck ./...
	exportloopref ./...
	golangci-lint run ./...
	gosec -conf ./gosec-config.json ./...
	yamllint -c ./yamllint-config.yaml ./cmd

## getversion: create txt file with version informations
getversion:
	date > ${BUILDINFO}
	go list >> ${BUILDINFO}
	git tag | tail -1 >> ${BUILDINFO}
	git log --oneline -n 1 >> $(BUILDINFO)
	go version >> $(BUILDINFO)
	uname -a >> $(BUILDINFO)

## compile: build binaries for all supported os
compile: getversion
	go build .

## build: cleanning and build new binaries
build: clean lint compile
	GOOS=linux go build -o ${BIN_LINUX}
	GOOS=darwin GOARCH=arm64 go build -o ${BIN_DARWIN_ARM}
	GOOS=darwin GOARCH=amd64 go build -o ${BIN_DARWIN_X8664}
	GOOS=windows GOARCH=amd64 go build -o ${BIN_WINDOWS}

## clean: delete compiled binaries
clean:
	rm -f ${BUILDINFO} \
		${APPNAME} \
		${BIN_DARWIN_ARM} \
		${BIN_DARWIN_X8664} \
		${BIN_LINUX} \
		${BIN_WINDOWS}
