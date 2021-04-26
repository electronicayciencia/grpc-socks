## Yet another socks proxy using gRPC streaming.

Cloned from elvizlai/grpc-socks to fix some points:

- Allow insecure gRPC
- Client death when DNS resolver fails
- Remove local dialing
- Server now runs in Alpine Linux and Raspberry Pi
- Use system DNS resolver (including /etc/hosts) if available

## Usage

First: Deploy server.

Launch client:

    ./exec_bin/client-linux -l 127.0.0.1:40000 -r grpcserver:9444 -d

If you are using curl:

    curl -x socks5://127.0.0.1:40000/ https://something.internal

For netcat:

    nc -X 5 -x 127.0.0.1:40000 192.168.451.23 443

Ssh (via nc):

    ssh -o ProxyCommand='nc -X 5 -x 127.0.0.1:40000 %h %p' user@internalserver.local




