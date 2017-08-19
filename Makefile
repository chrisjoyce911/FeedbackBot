.PHONY: all docker deps silent-test format test report

all: format kafkatohip.out bin/kafkatohip .git/hooks/pre-commit bin/kafka-console-producer

docker: silent-test bin/docker-kafkatohip bin/.docker-kafkatohip

kafkatohip.out: kafkatohip.go configmanager.go consumer.go hipchat.go messagemanager.go
	go test -coverprofile=bin/kafkatohip.out

bin/kafkatohip: kafkatohip.go configmanager.go consumer.go hipchat.go messagemanager.go
	go build -o bin/kafkatohip .

bin/docker-kafkatohip: kafkatohip.go configmanager.go consumer.go hipchat.go messagemanager.go
	SCGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/docker-kafkatohip .

bin/.docker-kafkatohip: bin/kafkatohip bin/docker-kafkatohip Dockerfile
	curl --remote-name --time-cond cacert.pem https://curl.haxx.se/ca/cacert.pem \
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