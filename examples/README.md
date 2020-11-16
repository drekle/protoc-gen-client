# Example protoc-gen-client

## Build this binary and ensure that it is in your `PATH`
Potentially you may just add the go bin directory to your path
```
export PATH=$PATH:$GOPATH/bin
```

In the project root
```
go install .
```

## Generate the example

From This directory:


This example extends several google protobuf message types.
```
git clone https://github.com/protocolbuffers/protobuf.git
```

Generate the command line client specified in `example.proto`.
```
protoc -I. -I.. -Iprotobuf/src/ --client_out=. example.proto
```