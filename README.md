# protobuf
Protobuf and gRPC helpers for working with Golang (for example an sql scan-able Timestamp well-known-type)

### How to use

```proto
syntax = "proto3";

package mypkg;

option go_package = "github.com/myname/mypkg";

import "veqryn/protobuf/timestamp.proto";

message MyMessage {
	veqryn.protobuf.Timestamp my_time = 1;
}
```

Then to compile, make sure to include the path to this repo like so:
```bash
protoc -I=/usr/local/include -I=${GOPATH}/src/github.com --go_out=plugins=grpc:${GOPATH}/src ${GOPATH}/src/github.com/myname/mypkg/mymessage.proto
```
