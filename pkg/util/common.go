package util

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func GetUserAgent(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	if x, ok := md["user-agent"]; ok {
		if x[0] != "" {
			return x[0]
		}
		return ""
	}
	return ""
}
