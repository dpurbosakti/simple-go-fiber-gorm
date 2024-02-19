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
