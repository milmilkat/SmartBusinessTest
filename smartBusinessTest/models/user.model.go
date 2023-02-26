package models

import (
	"time"

	"github.com/google/uuid"

)

type User struct {
	UserId    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FirstName string    `gorm:"type:varchar(255);not null"`
	LastName  string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"uniqueIndex;not null"`
	Created   time.Time
}

type AddUserRequest struct {
	FirstName string    `json:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
	Email     string    `json:"email,omitempty"`
}

type UpdateUser struct {
	FirstName string    `json:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
	Email     string    `json:"email,omitempty"`
}
