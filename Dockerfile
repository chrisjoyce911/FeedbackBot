FROM scratch
ADD bin/mydockerbot /
ENTRYPOINT ["/mydockerbot"]
CMD [""]
