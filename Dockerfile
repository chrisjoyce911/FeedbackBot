FROM golang:1.8
WORKDIR /usr/local/src/slacktohip/
COPY *.go /usr/local/src/slacktohip/
ENV GOPATH=/usr/local/
RUN echo $GOPATH
RUN go get -v ./... && GOOS=linux go build -tags netgo -installsuffix netgo -o docker-kafkatohip -v

#Need an image with x509 root certs. Can use scratch if you can download root certs from somewhere.
FROM scratch
COPY cert.pem /etc/ssl/certs/ca-certificates.crt
COPY --from=0 /usr/local/src/slacktohip/docker-kafkatohip .
ENTRYPOINT ["/docker-kafkatohip"]
CMD [""]