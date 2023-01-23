package models

import (
	"gorm.io/gorm"
)

type Office struct {
	gorm.Model
	Name   string
	Detail string

	Users []*User `gorm:"many2many:workspaces"`
}

type OfficesPaginate struct {
	OfficeList []*Office
	TotalCount uint32
}
