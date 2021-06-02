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

If `close: Bad file descriptor` try `nc -4`.

## Other ways to tunnel

### Tunnel with other netcats (no socks5 proxy):

Traditional:

    nc -l -p 8001 -c "nc 127.0.0.1 8000"

Busybox:

    nc -v -lk -p 8001 -e /usr/bin/nc 127.0.0.1 8000

### tsocks

### IPtables

Enable NAT:

    iptables -A PREROUTING -t nat -p tcp -i eth0 --dport 80 -j DNAT --to-destination 192.264.1.176:8000

Change the original IP to the intermediate box:

    iptables -A POSTROUTING -t nat -p tcp -d 192.264.1.176 --dport 8000 -j MASQUERADE

Allow packet forwarding:

    iptables -A FORWARD -p tcp -j ACCEPT

Enable kernel forwarding:

    sysctl net.ipv4.conf.eth0.forwarding=1

To delete iptables rules just change `-A` to `-D`.










