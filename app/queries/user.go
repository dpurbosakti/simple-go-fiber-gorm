package queries

import (
	"fmt"

	"github.com/dpurbosakti/fiber-gorm/app/models"
	"gorm.io/gorm"
)

type UserQuerier interface {
	CreateUser(u *models.User) error
	GetAllUser() ([]models.User, error)
	GetUserByID(ID string) (models.User, error)
	DeleteUser(ID string) error
	UpdateUser(u *models.User) error
}

func (q *Queries) CreateUser(u *models.User) error {
	if result := q.Debug().Create(u); result.Error != nil {
		return fmt.Errorf("failed to create user: %w", result.Error)
	}
	return nil
}

func (q *Queries) GetAllUser() ([]models.User, error) {
	users := []models.User{}
	if result := q.Debug().Find(&users); result.Error != nil {
		return nil, fmt.Errorf("failed to get all users data: %w", result.Error)
	}
	return users, nil
}

func (q *Queries) GetUserByID(ID string) (models.User, error) {
	user := models.User{}
	if result := q.Debug().Where("id = ?", ID).First(&user); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return user, fmt.Errorf("failed to get data user by id %s, %w", ID, result.Error)
		}
		return user, fmt.Errorf("failed to get all users data: %w", result.Error)
	}
	return user, nil
}

func (q *Queries) UpdateUser(u *models.User) error {
	if result := q.Debug().Save(u); result.Error != nil {
		return fmt.Errorf("failed to update user data: %w", result.Error)
	}

	return nil
}

func (q *Queries) DeleteUser(ID string) error {
	var user models.User
	if result := q.Debug().Where("id = ?", ID).Delete(&user, ID); result.Error != nil {
		return fmt.Errorf("failed to delete user: %w", result.Error)
	}

	return nil
}
