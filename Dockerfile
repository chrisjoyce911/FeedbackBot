FROM golang:1.8
WORKDIR /usr/local/src/slacktohip/
COPY *.go /usr/local/src/slacktohip/
ENV GOPATH=/usr/local/
RUN echo $GOPATH
RUN go get -v ./... && GOOS=linux go build -tags netgo -installsuffix netgo -o mydockerbot -v

FROM scratch
COPY --from=0 /usr/local/src/slacktohip/mydockerbot .
ENTRYPOINT ["/mydockerbot"]
CMD [""]
