FROM golang:1.10 AS fan-a-shitara-compile

RUN apt update && apt install -q -y unzip \
     --no-install-recommends \
     && rm -rf /var/lib/apt/lists/*
RUN wget https://github.com/google/protobuf/releases/download/v3.5.0/protoc-3.5.0-linux-x86_64.zip && unzip protoc-3.5.0-linux-x86_64.zip
RUN mv bin/protoc /usr/bin/ && chmod +x /usr/bin/protoc

ADD ./ /usr/src/test
WORKDIR /usr/src/test/src/
RUN go get -u google.golang.org/grpc \
              github.com/golang/protobuf/protoc-gen-go \
              github.com/BurntSushi/toml \
              github.com/gomodule/redigo/redis
RUN protoc --proto_path=. --go_out=plugins=grpc:./ protocol/protocol.proto
RUN cd /usr/src/test/src/server ; go build server.go
RUN cd /usr/src/test/src/client ; go build client.go

FROM ubuntu:18.04
COPY --from=fan-a-shitara-compile /usr/src/test/src/server/server /server
COPY --from=fan-a-shitara-compile /usr/src/test/src/client/client /client
ADD src/config /config
