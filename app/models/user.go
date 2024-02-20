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
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at" gorm:"default:now()"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"default:now()"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
}

type UpdateUserRequest struct {
	Name    *string `json:"name" validate:"required"`
	Email   *string `json:"email"`
	Address *string `json:"address"`
	Phone   *string `json:"phone"`
}

func CreateReqToEntity(req *CreateUserRequest) *User {
	return &User{
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
		Phone:   req.Phone,
	}
}

func UpdateReqToEntity(req *UpdateUserRequest, u *User) {
	if req.Name != nil {
		u.Name = *req.Name
	}

	if req.Email != nil {
		u.Email = *req.Email
	}

	if req.Address != nil {
		u.Address = *req.Address
	}

	if req.Phone != nil {
		u.Phone = *req.Phone
	}
}
