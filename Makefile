.PHONY: deps silent-test format test docker builddocker

all: bin/slacktohip bin/mydockerbot

bin/slacktohip: slacktohip.go slack.go
	go build -o bin/slacktohip .

bin/mydockerbot: slack.go slacktohip.go
	SCGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/mydockerbot .

builddocker:
	docker build -t slacktohip -f Dockerfile .

test:
	go test -v ./...

silent-test:
	go test ./...

format:
	go fmt ./...

deps:
	go get -v ./...
