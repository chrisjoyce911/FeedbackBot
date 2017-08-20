.PHONY: deps silent-test format test builddocker docker report

all: slacktohip.out bin/slacktohip bin/mydockerbot 

docker: silent-test bin/mydockerbot builddocker

bin/slacktohip: slacktohip.go slack.go
	go build -o bin/slacktohip .

bin/mydockerbot: slack.go slacktohip.go
	SCGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/mydockerbot .

slacktohip.out: slacktohip_test.go slack_test.go
	go test -coverprofile=bin/slacktohip.out

cacert.pem:
	curl --remote-name --time-cond cacert.pem https://curl.haxx.se/ca/cacert.pem

builddocker: cacert.pem
	docker build -t slacktohip -f Dockerfile .

test:
	go test -v -cover ./...

silent-test:
	go test ./...

format:
	go fmt ./...

deps:
	go get -v ./...

report:
	go tool cover -html=bin/slacktohip.out