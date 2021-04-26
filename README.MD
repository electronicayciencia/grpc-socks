## Yet another socks proxy using gRPC streaming.

Cloned from elvizlai/grpc-socks to solve some items:

- Allow insecure gRPC
- Fix client death if DNS resolver fails
- Remove local dialing
- Server now runs in Alpine Linux and Raspberry Pi




### Server

Using docker(using bin file if docker is not supported) to start a server.

```Dockerfile
docker run -itd \
    -p 10465:10465 \
    --name grpc-ss \
    --restart always \
    sdrzlyz/grpc-ss -l :10465
```

### Client

```
client_bin_file -r xxx.xxx.xxx.xxx:10465
```

#### You can download both server and client bin file from github release.
