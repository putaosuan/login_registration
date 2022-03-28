package util

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"net"
	"strconv"
	"strings"
)

// GetRealAddr get real client ip
func GetRealAddr(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "127.0.0.1"
	}
	fmt.Println("22222")
	rips := md.Get("x-real-ip")
	if len(rips) == 0 {
		return ""
	}

	return rips[0]
}

// GetPeerAddr get peer addr
func GetPeerAddr(ctx context.Context) string {
	var addr string
	if pr, ok := peer.FromContext(ctx); ok {
		if tcpAddr, ok := pr.Addr.(*net.TCPAddr); ok {
			addr = tcpAddr.IP.String()
		} else {
			addr = pr.Addr.String()
		}
	}
	return addr
}

//func GetClientIp(c *context.Context) string {
//	ip := c.ClientIP()
//	if ip == "::1" {
//		ip = "127.0.0.1"
//	}
//	return ip
//}
func IpStringToInt(ipstring string) int {
	if net.ParseIP(ipstring) == nil {
		return 0
	}
	ipSegs := strings.Split(ipstring, ".")
	var ipInt int = 0
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return ipInt
}
