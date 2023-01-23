package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Role string `gorm:"index"`
}

type RolesPaginate struct {
	RoleList   []*Role
	TotalCount uint32
}
