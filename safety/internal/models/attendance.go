package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Attendance struct {
	gorm.Model
	UserID           uuid.UUID `gorm:"primaryKey"`
	User             User      `gorm:"foreignKey:UserID;references:ID"`
	UserUsername     string    `gorm:"->;-:migration;"`
	UserAvatar       string    `gorm:"->;-:migration;"`
	ScheduleID       uint32    `gorm:"primaryKey"`
	Schedule         Schedule  `gorm:"foreignKey:ScheduleID;references:ID"`
	ScheduleDate     time.Time `gorm:"->;-:migration;"`
	OfficeID         string    `gorm:"->;-:migration;"`
	OfficeName       string    `gorm:"->;-:migration;"`
	ImageUrl         string
	Description      string
	AdminUsername    string
	Status           string `gorm:"default:pending;"`
	StatusAt         time.Time
	StatusInfo       string
	CheckTemperature float64
	CheckStatus      string    `gorm:"default:pending;"`
	CheckIn          time.Time `gorm:"default:null;"`
	CheckOut         time.Time `gorm:"default:null;"`
}

type AttendancesPaginate struct {
	AttendanceList []*Attendance
	TotalCount     uint32
}
