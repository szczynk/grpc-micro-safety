package models

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	OfficeID      uint32
	Office        Office `gorm:"foreignKey:OfficeID;references:ID"`
	OfficeName    string `gorm:"-:migration"`
	TotalCapacity uint32
	Capacity      uint32
	Date          time.Time `gorm:"uniqueIndex"`

	Users []*User `gorm:"many2many:attendances;"`
}

type SchedulesPaginate struct {
	ScheduleList []*Schedule
	TotalCount   uint32
}

type CreateSchedule struct {
	OfficeID      uint32
	TotalCapacity uint32
	Month         uint32
	Year          uint32
}

type ScheduleWithOffce struct {
	ID            uint
	OfficeID      uint32
	OfficeName    string
	TotalCapacity uint32
	Capacity      uint32
	Date          time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
