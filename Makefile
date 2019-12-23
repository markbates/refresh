TAGS ?= "sqlite"

install:
	packr2
	go install -v ./.

build:
	packr2
	go build -v .

test:
	packr2
	go test -tags ${TAGS} ./...

