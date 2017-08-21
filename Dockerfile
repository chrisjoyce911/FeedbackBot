FROM chrisjoyce911/goscratchcert
ADD hipchat.json /
ADD config.json /
ADD bin/docker-kafkatohip /
ENTRYPOINT ["/docker-kafkatohip"]
CMD [""]
