package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"
	"user/internal/domain"
	"user/internal/models"
	"user/pkg/grpc_errors"
	"user/pkg/logger"
	"user/pkg/utils"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
)

// User UseCase
type userUseCase struct {
	logger        logger.Logger
	userRepo      domain.UserRepository
	userRedisRepo domain.UserRedisRepository
}

// New User UseCase
func NewUserUseCase(logger logger.Logger, userRepo domain.UserRepository, userRedisRepo domain.UserRedisRepository) domain.UserUseCase {
	return &userUseCase{logger: logger, userRepo: userRepo, userRedisRepo: userRedisRepo}
}

// *Command

// create new user
func (u *userUseCase) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.CreateUser")
	defer span.Finish()

	newUser, err := u.userRepo.CreateUser(ctx, user)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "idx_users_email"):
			return nil, grpc_errors.ErrEmailExists
		case strings.Contains(err.Error(), "idx_users_username"):
			return nil, grpc_errors.ErrUsernameExists
		default:
			u.logger.Errorf("userRepo.CreateUser: %v", err)
			return nil, fmt.Errorf("userRepo.CreateUser: %v", err)
		}
	}

	return newUser, nil
}

// Login user with email and password
func (u *userUseCase) Login(ctx context.Context, email string, password string) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.Login")
	defer span.Finish()

	foundUser, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		u.logger.Errorf("userRepo.FindByEmail: %v", err)
		return nil, fmt.Errorf("userRepo.FindByEmail: %v", err)
	}

	if err := utils.CheckPasswordHash(foundUser.Password, password); err != nil {
		return nil, grpc_errors.ErrInvalidPassword
	}

	return foundUser, err
}

func (u *userUseCase) UpdateByID(ctx context.Context, ID uuid.UUID, updates models.User) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.UpdateByID")
	defer span.Finish()

	err := u.userRedisRepo.DeleteByID(ctx, "user_id:", ID.String())
	if err != nil {
		u.logger.Errorf("userRedisRepo.DeleteByID: %v", err)
	}

	updatedUser, err := u.userRepo.UpdateByID(ctx, ID, updates)
	if err != nil {
		u.logger.Errorf("userRepo.UpdateByID: %v", err)
		return nil, fmt.Errorf("userRepo.UpdateByID: %v", err)
	}

	return updatedUser, err
}

func (u *userUseCase) UpdateByEmail(ctx context.Context, email string, updates map[string]interface{}) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.UpdateByEmail")
	defer span.Finish()

	err := u.userRedisRepo.DeleteByID(ctx, "user_email", email)
	if err != nil {
		u.logger.Errorf("userRedisRepo.DeleteByID: %v", err)
	}

	user, err := u.userRepo.UpdateByEmail(ctx, email, updates)
	if err != nil {
		u.logger.Errorf("userRepo.UpdateByEmail: %v", err)
		return nil, fmt.Errorf("userRepo.UpdateByEmail: %v", err)
	}

	return user, err
}

func (u *userUseCase) DeleteByID(ctx context.Context, ID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.DeleteByID")
	defer span.Finish()

	err := u.userRedisRepo.DeleteByID(ctx, "user_id:", ID.String())
	if err != nil {
		u.logger.Errorf("userRedisRepo.DeleteByID: %v", err)
	}

	err = u.userRepo.DeleteByID(ctx, ID)
	if err != nil {
		u.logger.Errorf("userRepo.DeleteByID: %v", err)
		return fmt.Errorf("userRepo.DeleteByID: %v", err)
	}

	return nil
}

// *Query

func (u *userUseCase) Find(ctx context.Context, filters map[string]string, paginateQuery *utils.Pagination, expire time.Duration) ([]*models.User, uint32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.Find")
	defer span.Finish()

	keys := make([]string, 0, len(filters))
	for k := range filters {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var filterKey string
	parsedFilters := make(map[string]interface{}, len(filters))
	for _, k := range keys {
		if len(filters[k]) > 0 {
			filterKey += fmt.Sprintf("%v_%v-", k, filters[k])

			if k != "limit" && k != "page" && k != "sort" {
				parsedFilters[k] = filters[k]
			}
		}
	}
	filterKey = strings.TrimSuffix(filterKey, "-")

	var cachedUsers models.UsersPaginate
	cachedByte, er := u.userRedisRepo.FindByID(ctx, "user_list:", filterKey)
	if er != nil {
		foundUserList, totalCount, err := u.userRepo.Find(ctx, parsedFilters, paginateQuery)
		if err != nil {
			u.logger.Errorf("userRepo.Find: %v", err)
			return nil, 0, fmt.Errorf("userRepo.Find: %v", err)
		}

		foundUsers := models.UsersPaginate{
			UserList:   foundUserList,
			TotalCount: totalCount,
		}

		foundUserByte, err := json.Marshal(foundUsers)
		if err != nil {
			u.logger.Errorf("json.Marshal: %v", err)
			return nil, 0, err
		}

		err = u.userRedisRepo.CreateUser(ctx, "user_list:", filterKey, foundUserByte, expire)
		if err != nil {
			u.logger.Errorf("userRedisRepo.CreateUser", err)
		}

		return foundUserList, totalCount, nil
	}

	err := json.Unmarshal(cachedByte, &cachedUsers)
	if err != nil {
		u.logger.Errorf("json.Unmarshal: %v", err)
		return nil, 0, err
	}

	return cachedUsers.UserList, cachedUsers.TotalCount, nil
}

func (u *userUseCase) FindByID(ctx context.Context, ID uuid.UUID, expire time.Duration) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.FindByID")
	defer span.Finish()

	cachedUser := new(models.User)
	cachedByte, er := u.userRedisRepo.FindByID(ctx, "user_id:", ID.String())
	if er != nil {
		foundUser, err := u.userRepo.FindByID(ctx, ID)
		if err != nil {
			u.logger.Errorf("userRepo.FindByID: %v", err)
			return nil, fmt.Errorf("userRepo.FindByID: %v", err)
		}

		foundUserByte, err := json.Marshal(foundUser)
		if err != nil {
			u.logger.Errorf("json.Marshal: %v", err)
			return nil, err
		}

		err = u.userRedisRepo.CreateUser(ctx, "user_id:", ID.String(), foundUserByte, expire)
		if err != nil {
			u.logger.Errorf("userRedisRepo.CreateUser: %v", err)
		}

		return foundUser, nil
	}

	err := json.Unmarshal(cachedByte, &cachedUser)
	if err != nil {
		u.logger.Errorf("json.Unmarshal: %v", err)
		return nil, err
	}

	return cachedUser, nil
}

func (u *userUseCase) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.FindByEmail")
	defer span.Finish()

	user, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		u.logger.Errorf("userRepo.FindByEmail: %v", err)
		return nil, fmt.Errorf("userRepo.FindByEmail: %v", err)
	}

	return user, err
}

func (u *userUseCase) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.FindByUsername")
	defer span.Finish()

	user, err := u.userRepo.FindByUsername(ctx, username)
	if err != nil {
		u.logger.Errorf("userRepo.FindByUsername: %v", err)
		return nil, fmt.Errorf("userRepo.FindByUsername: %v", err)
	}

	return user, err
}

func (u *userUseCase) FindByVerificationCode(ctx context.Context, code string) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.FindByVerificationCode")
	defer span.Finish()

	user, err := u.userRepo.FindByVerificationCode(ctx, code)
	if err != nil {
		u.logger.Errorf("userRepo.FindByVerificationCode: %v", err)
		return nil, fmt.Errorf("userRepo.FindByVerificationCode: %v", err)
	}

	return user, err
}

func (u *userUseCase) FindByResetPasswordToken(ctx context.Context, code string) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserUseCase.FindByResetPasswordToken")
	defer span.Finish()

	user, err := u.userRepo.FindByResetPasswordToken(ctx, code)
	if err != nil {
		u.logger.Errorf("userRepo.FindByResetPasswordToken: %v", err)
		return nil, fmt.Errorf("userRepo.FindResetPasswordToken: %v", err)
	}

	return user, err
}
