package repositories

import (
	"Gintest/src/entities"
	"Gintest/src/managers"
	"errors"
	"gorm.io/gorm"
)

type UserRepositoryContract interface {
	Find(int64) (*entities.User, error)
	FindByEmail(string) (*entities.User, error)
	Create(*entities.User) (*entities.User, error)
	Update(*entities.User, map[string]interface{}) (*entities.User, error)
}

type UserRepository struct {
	DatabaseManager managers.DatabaseManagerContract
}

func (repository *UserRepository) Find(id int64) (*entities.User, error) {
	user := entities.User{}

	result := repository.DatabaseManager.Getclient().Where("id = ?", id).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &user, result.Error
	}

	if result.Error != nil {
		panic(result.Error)
	}

	return &user, nil
}

func (repository *UserRepository) FindByEmail(email string) (*entities.User, error) {

	user := entities.User{}

	result := repository.DatabaseManager.Getclient().Where("email = ?", email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &user, result.Error
	}

	if result.Error != nil {
		panic(result.Error)
	}

	return &user, nil
}

func (repository *UserRepository) Create(user *entities.User) (*entities.User, error) {

	result := repository.DatabaseManager.Getclient().Create(user)

	return user, result.Error
}

func (repository *UserRepository) Update(user *entities.User, data map[string]interface{}) (*entities.User, error) {

	result := repository.DatabaseManager.Getclient().Model(user).Updates(data)

	return user, result.Error
}
