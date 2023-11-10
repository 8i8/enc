all:
	go run cmd/env.go

build:
	go build -o enc cmd/env.go

install:
	go build -o enc cmd/env.go
	mv enc /usr/local/bin

clean:
	- rm enc
	- rm enc.svg

draw:
	- rm notes/enc.svg
	plantuml -tsvg notes/enc.puml
	firefox-esr notes/enc.svg
