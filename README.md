# Protocol Buffers, Go, and RPC

This is an example project that shows how to set up and use protocol buffers
and RPC with Go.

## Installing the Procotol Buffers compiler, `protoc`

To install the base protocol buffers compiler:

``` bash
cd /tmp
curl https://protobuf.googlecode.com/files/protobuf-2.5.0.tar.bz2 > protobuf-2.5.0.tar.bz2
tar -xf protobuf-2.5.0.tar.bz2
cd protobuf-2.5.0
./configure
make
sudo make install
```

*NOTE:* This has only been tested on OSX, but should work with any Unix that
has GCC installed.

To install the Go generator for protoc:

``` bash
go get -v -u github.com/kylelemons/go-rpcgen/protoc-gen-go
```

And we're done!

## Generate the Go code services

Note that you'll have to do this step before this code will build.

``` bash
protoc --go_out=src/counters counters.proto
```

It should return silently if nothing bad happened!
