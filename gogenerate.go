package protobuf

// To run from project root directory: `go generate -x ./...`

// Generate golang
//go:generate protoc -I=/usr/local/include -I=${GOPATH}/src/github.com --go_out=plugins=grpc:${GOPATH}/src ${GOPATH}/src/github.com/veqryn/protobuf/timestamp.proto

// Generate test golang
//go:generate protoc -I=/usr/local/include -I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I=${GOPATH}/src/github.com/envoyproxy -I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway -I=${GOPATH}/src/github.com --go_out=plugins=grpc:${GOPATH}/src --grpc-gateway_out=logtostderr=true:${GOPATH}/src --swagger_out=logtostderr=true:${GOPATH}/src/github.com --validate_out=lang=go:${GOPATH}/src ${GOPATH}/src/github.com/veqryn/protobuf/ptypes/test/test.proto
