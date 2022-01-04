all:
	go run cmd/cli/cli.go

build:
	go build -o enc cmd/cli/cli.go

install:
	go build -o enc cmd/cli/cli.go
	mv enc ~/.bin

clean:
	- rm enc
	- rm enc.svg

draw:
	- rm notes/enc.svg
	plantuml -tsvg notes/enc.puml
	firefox-esr notes/enc.svg
