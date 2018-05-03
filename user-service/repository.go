// user-service/repository.go

package main

import (
	userProto "github.com/fidelisojeah/go-microservice/user-service/proto/auth"
	"github.com/jinzhu/gorm"
)

// Repository class
type Repository interface {
	Create(user *userProto.User) error
	Get(id string) (*userProto.User, error)
	GetAll() ([]*userProto.User, error)
	GetByEmailandPassword(user *userProto.User) (*userProto.User, error)
	GetByEmail(emailAddress string) (*userProto.User, error)
	// Validate (*userProto.Token)
	// Close()
}

// UserRepository - Class
type UserRepository struct {
	db *gorm.DB
}

// Create -  Creates a user in the dB
func (repo *UserRepository) Create(user *userProto.User) error {
	err := repo.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

// Get - retrieves user from db Based on the id
func (repo *UserRepository) Get(id string) (*userProto.User, error) {
	var user *userProto.User
	user.Id = id
	err := repo.db.Select("Name, Email, Company").First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetAll - retrieves all users from db
func (repo *UserRepository) GetAll() ([]*userProto.User, error) {
	var users []*userProto.User
	err := repo.db.Select("Name, Email, Company").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetByEmailandPassword - retrieves from database username and password
func (repo *UserRepository) GetByEmailandPassword(user *userProto.User) (*userProto.User, error) {
	err := repo.db.First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetByEmail - retrieves from database username and password
func (repo *UserRepository) GetByEmail(emailAddress string) (*userProto.User, error) {
	user := &userProto.User{}

	err := repo.db.Where("email = ?", emailAddress).First(&user).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}
