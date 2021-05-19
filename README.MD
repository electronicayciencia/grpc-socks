## Yet another socks proxy using gRPC streaming.

Cloned from elvizlai/grpc-socks to fix some points:

- Allow insecure gRPC
- Prevent client death when DNS resolver fails
- Remove local dialing
- Server now runs in Alpine Linux and Raspberry Pi
- Use system DNS resolver (including /etc/hosts) if available

## Build

1. Run make to build binaries
1. Build the docker image

       docker build . -t grpcserver:latest

## Usage

First: Deploy server.

And launch client:

    ./exec_bin/client-linux -l 127.0.0.1:40000 -r grpcserver:9444 -d

## Common applications

### If you are using curl

    curl -x socks5://127.0.0.1:40000/ https://something.internal

### For netcat (openbsd version)

    nc -X 5 -x 127.0.0.1:40000 192.168.451.23 443

### Ssh (via nc)

    ssh -o ProxyCommand='nc -X 5 -x 127.0.0.1:40000 %h %p' user@internalserver.local

Useful options:

	ssh \
	  -o ProxyCommand='nc -X 5 -x 127.0.0.1:40000 %h %p' \
	  -o StrictHostKeyChecking=no \
	  -o UserKnownHostsFile=/dev/null \
	  user@internalserver.local

### To tunnel remote ports:

    mkfifo fifo
    nc -l 636 0<fifo | nc -X 5 -x 127.0.0.1:40000 internalserver.local 636 1>fifo

### Tunnel with other netcats (no socks5 proxy):

Traditional:

    nc -l -p 8001 -c "nc 127.0.0.1 8000"

Busybox:

    nc -v -lk -p 8001 -e /usr/bin/nc 127.0.0.1 8000


Use tsocks for other commands.


