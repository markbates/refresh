install:
	go install -v .

test:
	go test -cover -failfast ./...