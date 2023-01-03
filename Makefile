cat:
	cat Makefile
.PHONY: cat

run:
	go run .

build:
	go build .

small:
	go build -ldflags="-s -w" .

fmt:
	go fmt ./...

lint:
	golint ./...

vet:
	go vet ./...

g-lint:
	golangci-lint run

strip:
	strip -s alap

install:
	go install .

clean:
	rm alap
