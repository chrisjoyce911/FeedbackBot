PACKAGES := github.com/chrisjoyce911/slacktohip
DEPENDENCIES := github.com/andybons/hipchat

all: build silent-test docker

docker:
	SCGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/mydockerbot .

builddocker:
	docker build -t slacktohip -f Dockerfile .

build:
	go build -o bin/slacktohip .

test:
	go test -v $(PACKAGES)

silent-test:
	go test $(PACKAGES)

format:
	go fmt $(PACKAGES)

deps:
	go get $(DEPENDENCIES)