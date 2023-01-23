package models

import "gorm.io/gorm"

type Method struct {
	gorm.Model
	Method string `gorm:"index"`
}

type MethodsPaginate struct {
	MethodList []*Method
	TotalCount uint32
}
