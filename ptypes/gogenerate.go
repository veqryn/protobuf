package ptypes

// To run from project root directory: `go generate -x ./ptypes/...`

// Generate golang protobuf/grpc/gateway and swagger docs
//go:generate protoc -I=/usr/local/include -I=./timestamp --go_out=plugins=grpc:${GOPATH}/src ./timestamp/timestamp.proto
