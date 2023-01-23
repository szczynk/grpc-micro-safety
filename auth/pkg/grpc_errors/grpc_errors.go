package grpc_errors

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-redis/redis/v9"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

var (
	ErrNotFound           = errors.New("Not found")
	ErrNoCtxMetaData      = errors.New("No ctx metadata")
	ErrNoAccessToken      = errors.New("Access token doesn't exist")
	ErrNoRefreshToken     = errors.New("Refresh token doesn't exist")
	ErrNoUserId           = errors.New("User ID doesn't exist")
	ErrNoPassword         = errors.New("No password input")
	ErrInvalidAuthHeader  = errors.New("Invalid authorization header")
	ErrInvalidAccessToken = errors.New("Invalid access token")
	ErrInvalidSessionId   = errors.New("Invalid session id")
	ErrInvalidEmail       = errors.New("Invalid email")
	ErrInvalidPassword    = errors.New("Invalid Password")
	ErrMethodNotAllowed   = errors.New("Method not allowed")
	ErrEmailExists        = errors.New("Email already exists")
	ErrUsernameExists     = errors.New("Username already exists")
	ErrNoEmailVerifyCode  = errors.New("Email verification code doesn't exist")
	ErrNoResetPassToken   = errors.New("Reset password token doesn't exist")
)

// Parse error and get code
func ParseGRPCErrStatusCode(err error) codes.Code {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return codes.NotFound
	case errors.Is(err, redis.Nil):
		return codes.NotFound
	case strings.Contains(err.Error(), "redis"):
		return codes.NotFound
	case errors.Is(err, ErrNotFound):
		return codes.NotFound

	case errors.Is(err, context.Canceled):
		return codes.Canceled

	case errors.Is(err, context.DeadlineExceeded):
		return codes.DeadlineExceeded

	case errors.Is(err, ErrEmailExists):
		return codes.AlreadyExists
	case errors.Is(err, ErrUsernameExists):
		return codes.AlreadyExists

	case errors.Is(err, ErrNoCtxMetaData):
		return codes.Unauthenticated
	case errors.Is(err, ErrNoAccessToken):
		return codes.Unauthenticated
	case errors.Is(err, ErrNoRefreshToken):
		return codes.Unauthenticated
	case errors.Is(err, ErrNoUserId):
		return codes.Unauthenticated

	case errors.Is(err, ErrInvalidSessionId):
		return codes.PermissionDenied
	case errors.Is(err, ErrInvalidAuthHeader):
		return codes.PermissionDenied
	case errors.Is(err, ErrInvalidAccessToken):
		return codes.PermissionDenied
	case errors.Is(err, ErrMethodNotAllowed):
		return codes.PermissionDenied

	case strings.Contains(err.Error(), "Validate"):
		return codes.InvalidArgument
	case errors.Is(err, ErrNoEmailVerifyCode):
		return codes.InvalidArgument
	case errors.Is(err, ErrNoResetPassToken):
		return codes.InvalidArgument
	case errors.Is(err, ErrInvalidEmail):
		return codes.InvalidArgument
	case errors.Is(err, ErrInvalidPassword):
		return codes.InvalidArgument
	case errors.Is(err, ErrNoPassword):
		return codes.InvalidArgument

	}
	return codes.Internal
}

// Map GRPC errors codes to http status
func MapGRPCErrCodeToHttpStatus(code codes.Code) int {
	switch code {
	case codes.Unauthenticated:
		return http.StatusUnauthorized

	case codes.AlreadyExists:
		return http.StatusBadRequest

	case codes.NotFound:
		return http.StatusNotFound

	case codes.Internal:
		return http.StatusInternalServerError

	case codes.PermissionDenied:
		return http.StatusForbidden

	case codes.Canceled:
		return http.StatusRequestTimeout

	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout

	case codes.InvalidArgument:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}
