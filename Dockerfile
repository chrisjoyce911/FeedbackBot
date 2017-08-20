FROM scratch
ADD bin/ca-bundle.crt /etc/ssl/certs/ca-certificates.crt
ADD hipchat.json /
ADD config.json /
ADD bin/docker-kafkatohip /
ENTRYPOINT ["/docker-kafkatohip"]
CMD [""]
