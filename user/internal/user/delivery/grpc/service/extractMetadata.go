package service

import (
	"context"
	"user/pkg/grpc_errors"

	"google.golang.org/grpc/metadata"
)

// extract user_id from metadata
func (u *usersService) ExtractUserId(ctx context.Context) (userId string, userRole string, err error) {
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
