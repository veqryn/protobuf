package protobuf

// To run from project root directory: `go generate -x ./...`

// Generate golang
//go:generate protoc -I=/usr/local/include -I=${GOPATH}/src/github.com/veqryn/protobuf --go_out=plugins=grpc:${GOPATH}/src ${GOPATH}/src/github.com/veqryn/protobuf/timestamp.proto
