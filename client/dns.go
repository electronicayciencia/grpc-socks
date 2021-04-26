package main

import (
	"context"
	"net"
	"log"

	"github.com/coocood/freecache"

	"github.com/elvizlai/grpc-socks/pb"
)

type DNSResolver struct {
	cache *freecache.Cache
}

var expireSeconds = 7200

var nameCtxKey = struct{}{}

// DNSResolver uses the remote DNS to resolve host names
func (d DNSResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	ctx = context.WithValue(ctx, nameCtxKey, name)

	log.Println("[DNS] resolve:", name)

	if v, err := d.cache.Get([]byte(name)); err == nil {
		log.Println("[DNS]", name, "cached")
		return ctx, v, nil
	}

	addr, err := net.ResolveIPAddr("ip", name)
	if err == nil {
		log.Println("[DNS]", name, "resolved using local:", addr.IP)
		d.cache.Set([]byte(name), addr.IP, expireSeconds)
		return ctx, addr.IP, err
	}

	ipResp, err := proxyClient.ResolveIP(ctx, &pb.IPAddr{
		Address: name,
	})

	if err == nil {
		log.Println("[DNS]", name, "resolved using remote:", ipResp.Data)
		d.cache.Set([]byte(name), ipResp.Data, expireSeconds)
		return ctx, ipResp.Data, err
	}

	log.Println("[DNS] Unable to resolve", name)
	return ctx, nil, err
}
