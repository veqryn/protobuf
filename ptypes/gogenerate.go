package ptypes

// To run from project root directory: `go generate -x ./ptypes/...`

// Generate golang protobuf/grpc/gateway and swagger docs
//go:generate protoc -I=/usr/local/include -I=${GOPATH}/src/github.com/veqryn/protobuf/ptypes/timestamp --go_out=plugins=grpc:${GOPATH}/src ${GOPATH}/src/github.com/veqryn/protobuf/ptypes/timestamp/timestamp.proto
