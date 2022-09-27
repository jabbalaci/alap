cat:
	cat Makefile

run:
	go run .

build:
	go build .

strip:
	strip -s alap

install:
	go install .

clean:
	rm alap
