.PHONY: all docker deps silent-test format test report clean
SHELL := /bin/sh

all: format bin/kafkatohip.out bin/kafkatohip .git/hooks/pre-commit bin/kafka-console-producer

help: ## This help message
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/:.*##\s*/:/' -e 's/^\(.\+\):\(.*\)/\\x1b[36m\1\\x1b[m:\2/ '

docker: silent-test bin/docker-kafkatohip bin/.docker-kafkatohip ## Builds Docker image

bin/kafkatohip.out: kafkatohip.go configmanager.go consumer.go hipchat.go messagemanager.go
	go test -coverprofile=bin/kafkatohip.out

bin/kafkatohip: kafkatohip.go configmanager.go consumer.go hipchat.go messagemanager.go
	go build -o bin/kafkatohip .

bin/docker-kafkatohip: kafkatohip.go configmanager.go consumer.go hipchat.go messagemanager.go
	SCGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/docker-kafkatohip .

bin/.docker-kafkatohip: bin/kafkatohip bin/docker-kafkatohip Dockerfile
	docker build -t kafkatohip -f Dockerfile .
	touch bin/.docker-kafkatohip

bin/kafka-console-producer: kafka-console-producer/kafka-console-producer.go
	go build -o bin/kafka-console-producer kafka-console-producer/kafka-console-producer.go

.git/hooks/pre-commit: pre-commit
	ln -s ../../pre-commit .git/hooks/pre-commit

test:
	go test -v -cover ./...

silent-test:
	go test ./...

format:
	go fmt ./...

deps:
	go get -v ./...

report:
	go tool cover -html=bin/kafkatohip.out

clean:
	-rm bin/.docker-kafkatohip
	-rm bin/kafka-console-producer
	-rm bin/docker-kafkatohip
	-rm bin/kafkatohip.out
	-rm bin/kafkatohip
