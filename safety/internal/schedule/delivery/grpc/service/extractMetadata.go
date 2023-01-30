package service

import (
	"context"
	"safety/pkg/grpc_errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (u *schedulesService) SendHeader(ctx context.Context) error {
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
