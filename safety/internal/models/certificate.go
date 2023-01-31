package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// join only get certificate lists
type Certificate struct {
	gorm.Model
	UserID        uuid.UUID
	User          User   `gorm:"foreignKey:UserID;references:ID"`
	UserUsername  string `gorm:"->;-:migration;"`
	UserAvatar    string `gorm:"->;-:migration;"`
	Dose          uint32
	ImageUrl      string
	Description   string
	AdminUsername string
	Status        string `gorm:"default:pending;"`
	StatusAt      time.Time
	StatusInfo    string
}

type CertificatesPaginate struct {
	CertificateList []*Certificate
	TotalCount      uint32
}
