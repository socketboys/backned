FROM ubuntu:latest
LABEL authors="rajatkr"

RUN apt-get update -qq

ENTRYPOINT ["top", "-b"]