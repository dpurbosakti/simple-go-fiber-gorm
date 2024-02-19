package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"autoIncrement;primaryKey"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Address   string         `json:"address"`
	Phone     string         `json:"phone"`
	CreatedAt time.Time      `json:"created_at" gorm:"default:now()"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"default:now()"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type CreateUserRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required"`
	Address string `json:"address" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
}

func ReqToEntity(req *CreateUserRequest) *User {
	return &User{
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
		Phone:   req.Phone,
	}
}
