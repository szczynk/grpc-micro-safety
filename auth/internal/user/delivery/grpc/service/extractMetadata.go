package service

import (
	"auth/pkg/grpc_errors"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader            = "user-agent"
	xForwardedForHeader        = "x-forwarded-for"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
}

func (u *usersService) ExtractMetadata(ctx context.Context) (*Metadata, error) {
	mtdt := &Metadata{}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc_errors.ErrNoCtxMetaData
	}

	if userAgents := md.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
		mtdt.UserAgent = userAgents[0]
	}

	if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
		mtdt.UserAgent = userAgents[0]
	}

	if mdClientIPs := md.Get(xForwardedForHeader); len(mdClientIPs) > 0 {
		mtdt.ClientIP = mdClientIPs[0]
	}

	if mdClientIPs := md.Get("x-real-ip"); len(mdClientIPs) > 0 {
		mtdt.ClientIP = mdClientIPs[0]
	}

	if p, ok := peer.FromContext(ctx); ok {
		mtdt.ClientIP = p.Addr.String()
	}

	if len(mtdt.ClientIP) == 0 {
		mtdt.ClientIP = "unknown"
	}

	return mtdt, nil
}

func (u *usersService) SendHeader(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return grpc_errors.ErrNoCtxMetaData
	}

	rateLimit := md.Get("X-RateLimit-Limit")
	rateRemaining := md.Get("X-RateLimit-Remaining")
	rateReset := md.Get("X-RateLimit-Reset")

	header := metadata.New(
		map[string]string{
			"X-RateLimit-Limit":     rateLimit[0],
			"X-RateLimit-Remaining": rateRemaining[0],
			"X-RateLimit-Reset":     rateReset[0],
		},
	)

	err := grpc.SendHeader(ctx, header)
	if err != nil {
		return status.Errorf(codes.Internal, "unable to send grpc header")
	}

	return nil
}
