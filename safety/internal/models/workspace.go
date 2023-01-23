package models

import (
	"github.com/google/uuid"
)

type Workspace struct {
	UserID   uuid.UUID
	OfficeID uint32
}

type WorkspacesPaginate struct {
	WorkspaceList []*User
	TotalCount    uint32
}
