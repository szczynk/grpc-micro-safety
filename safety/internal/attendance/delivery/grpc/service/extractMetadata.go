package service

import (
	"context"
	"safety/pkg/grpc_errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// extract user_id from metadata
func (u *attendancesService) ExtractUserId(ctx context.Context) (userId string, userRole string, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", "", grpc_errors.ErrNoCtxMetaData
	}

	userID := md.Get("user_id")
	if len(userID) == 0 {
		return "", "", grpc_errors.ErrNoUserId
	}

	userrole := md.Get("role")
	if len(userrole) == 0 {
		return "", "", grpc_errors.ErrNoUserId
	}

	userId, userRole = userID[0], userrole[0]
	return userId, userRole, nil
}

func (u *attendancesService) SendHeader(ctx context.Context) error {
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
