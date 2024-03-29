FROM golang:1.12.9-buster

# Install requirements for Golang
RUN set -eux; \
  apt-get update -y; \
  apt-get upgrade -y; \
  apt-get install -y unzip wget; \
  wget -O /tmp/protoc.zip "https://github.com/protocolbuffers/protobuf/releases/download/v3.9.1/protoc-3.9.1-linux-x86_64.zip"; \
  unzip /tmp/protoc.zip -d /tmp/; \
  cp -ravp /tmp/include/ /usr/local/; \
  cp -ravp /tmp/bin/ /usr/; \
  go get -u github.com/grpc-ecosystem/grpc-gateway/...; \
  go get -u github.com/envoyproxy/protoc-gen-validate; \
  go get -u google.golang.org/grpc; \
  go get -u github.com/golang/mock/...; \
  go build -i -o "$(go env GOPATH)"/bin/cover cmd/cover; \
  go get -u github.com/kisielk/errcheck; \
  go get -d -u github.com/golang/protobuf/protoc-gen-go; \
  git -C "$(go env GOPATH)"/src/github.com/golang/protobuf checkout "v1.3.2"; \
  go install github.com/golang/protobuf/protoc-gen-go; \
  rm -rf /tmp/* /var/tmp/* /var/cache/yum;

ENTRYPOINT ["go"]
CMD ["generate", "-x", "github.com/..."]
