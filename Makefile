default: test

test:
	cd calculator && go build -v ./...

build:
	cd cmd/calculator && go test -v ./...