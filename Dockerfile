FROM golang:1.8
WORKDIR /usr/local/src/slacktohip/
COPY *.go /usr/local/src/slacktohip/
ENV GOPATH=/usr/local/
RUN echo $GOPATH
RUN go get -v ./... && GOOS=linux go build -tags netgo -installsuffix netgo -o mydockerbot -v

#Need an image with x509 root certs. Can use scratch if you can download root certs from somewhere.
FROM scratch
COPY cacert.pem /etc/ssl/certs/ca-certificates.crt
COPY --from=0 /usr/local/src/slacktohip/mydockerbot .
ENTRYPOINT ["/mydockerbot"]
CMD [""]
