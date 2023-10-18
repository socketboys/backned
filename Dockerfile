FROM ubuntu:latest
LABEL authors="rajatkr"

ENV GO_VERSION=1.21.0
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

RUN apt-get update
RUN apt-get install -y wget git gcc

RUN wget -P /tmp "https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz"

RUN tar -C /usr/local -xzf "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"
RUN rm "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR $GOPATH

WORKDIR /tmp

RUN apt-get update

RUN apt-get install -y python3-pip
RUN apt-get install -y python3-scipy

COPY /pipeline-cli/requirements.txt .

RUN pip3 install -r requirements.txt

RUN mkdir /backned
WORKDIR /tmp/backned

ADD . .

ENTRYPOINT ["go", "run", "cmd/main.go"]