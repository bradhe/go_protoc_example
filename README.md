# RPC in Go with Protocol Buffers

This is an example project that shows how to set up and use protocol buffers
and RPC with Go.

## Layout

There are three folders under `src`:

1. `counters` is where the library code lives, including the generated protoc
   code.
1. `counters_server` is the server component
1. `counters_client` is a CLI for interacting with the server.

## Getting Up and Running

### 1. Install Google's Procotol Buffers compiler, `protoc`

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

**NOTE:** This has only been tested on OSX, but should work with any Unix that
has GCC installed.

To install the Go generator for protoc:

``` bash
go get -v -u github.com/kylelemons/go-rpcgen/protoc-gen-go
```

And we're done!

### 2. Generate the Go code

Note that you'll have to do this step before this code will build.

``` bash
protoc --go_out=src/counters counters.proto
```

It should return silently if nothing bad happened!

### 3. Build and run the client and server

Pretty straight forward:

``` bash
go build counters_client
go build counters_server
```

If your `GOPATH` isn't set yet:

``` bash
export GOPATH=`pwd`
```

#### Start the Server

``` bash
./counters_server
```

#### Run the client

``` bash
$ ./counters_client inc "Hello, World"
Hello, World: 1
```

## Feedback

I'd really love to hear your thoughts. How can I make this more idiomatic? Are there better patterns for this or that? Here's how you can help:

1. Fork this repo.
1. Commit any fixes/changes/etc.
1. If you are so compelled, send a pull request.
1. Send me an email, if you want: brad.heller@gmail.com
