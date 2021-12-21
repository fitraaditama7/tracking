download:
	go mod download

run:
	go run *.go

all:
	go mod download
	go run *.go

.PHONY: download run