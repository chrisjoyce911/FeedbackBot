.PHONY: all docker deps silent-test format test report

all: format kafkatohip.out bin/kafkatohip

docker: silent-test bin/docker-kafkatohip bin/.docker-kafkatohip

kafkatohip.out: kafkatohip_test.go readconfig.go
	go test -coverprofile=bin/kafkatohip.out

bin/kafkatohip: kafkatohip_test.go readconfig.go
	go build -o bin/kafkatohip .

bin/docker-kafkatohip: kafkatohip.go readconfig.go
	SCGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/docker-kafkatohip .

bin/.docker-kafkatohip: kafkatohip bin/docker-kafkatohip Dockerfile
	curl --remote-name --time-cond cacert.pem https://curl.haxx.se/ca/cacert.pem \
	docker build -t kafkatohip -f Dockerfile .
	touch bin/.docker-kafkatohip

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