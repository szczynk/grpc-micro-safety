package models

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Service string `gorm:"index"`
}

type ServicesPaginate struct {
	ServiceList []*Service
	TotalCount  uint32
}
