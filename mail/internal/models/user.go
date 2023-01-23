package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                 uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username           string    `gorm:"uniqueIndex;not null"`
	Email              string    `gorm:"uniqueIndex;not null"`
	Password           string    `gorm:"not null"`
	Role               string    `gorm:"type:varchar(255);not null"`
	Avatar             string
	VerificationCode   string `gorm:"index"`
	Verified           bool   `gorm:"default:false;"`
	VerifiedAt         time.Time
	PasswordResetToken string `gorm:"index"`
	PasswordResetAt    time.Time
	Otp_enabled        bool `gorm:"default:false;"`
	Otp_verified       bool `gorm:"default:false;"`
	Otp_secret         string
	Otp_auth_url       string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *gorm.DeletedAt
}
