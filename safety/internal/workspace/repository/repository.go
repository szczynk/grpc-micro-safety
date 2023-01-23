package repository

import (
	"context"
	"fmt"
	"safety/internal/domain"
	"safety/internal/models"
	"safety/pkg/utils"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

// Workspace repository
type workspaceRepo struct {
	db *gorm.DB
}

func NewWorkspaceRepository(db *gorm.DB) domain.WorkspaceRepository {
	return &workspaceRepo{db: db}
}

// *Command

func (ur *workspaceRepo) CreateWorkspace(ctx context.Context, workspace *models.Workspace) (*models.Workspace, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "WorkspaceRepo.CreateWorkspace")
	defer span.Finish()

	err := ur.db.WithContext(ctx).Create(&workspace).Error
	if err != nil {
		return nil, err
	}

	return workspace, nil
}

func (ur *workspaceRepo) DeleteByUserID(ctx context.Context, UserID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "WorkspaceRepo.DeleteByUserID")
	defer span.Finish()

	workspace := new(models.Workspace)
	err := ur.db.WithContext(ctx).
		Where("user_id = ?", UserID.String()).
		Delete(&workspace).Error
	if err != nil {
		return err
	}

	return nil
}

// *Query

func (ur *workspaceRepo) Find(ctx context.Context, filter map[string]interface{}, paginateQuery *utils.Pagination) ([]*models.User, uint32, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "WorkspaceRepo.Find")
	defer span.Finish()

	var users []*models.User
	var count int64

	result := ur.db.WithContext(ctx).Table("workspaces").
		Select("users.id, users.email, users.username, users.role, users.avatar, users.verified, users.created_at, users.updated_at").
		Joins("JOIN users ON users.id=workspaces.user_id")
	for k, v := range filter {
		switch k {
		case "verified":
			result = result.Where(k+" = ?", v)
		case "office_id":
			result = result.Where(k+" = ?", v)
		default:
			v1 := fmt.Sprintf("%%%v%%", v)
			result = result.Where(k+" LIKE ?", v1)
		}
	}

	result.Count(&count)

	rows := result.Offset(int(paginateQuery.GetOffset())).
		Limit(int(paginateQuery.GetLimit())).
		Order(paginateQuery.GetSort()).
		Scan(&users)
	if rows.Error != nil {
		return nil, 0, rows.Error
	}

	return users, uint32(count), nil
}
