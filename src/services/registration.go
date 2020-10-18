package services

import (
	"Gintest/src/entities"
	"Gintest/src/repositories"
	"Gintest/src/utilities"
	"errors"
)

type RegistrationServiceContract interface {
	Register(*entities.UserRegisteration) (*entities.User, error)
}

type RegistrationService struct {
	UserRepository repositories.UserRepositoryContract
}

func (service *RegistrationService) Register(userRegisteration *entities.UserRegisteration) (*entities.User, error) {

	var err error
	var user *entities.User

	user, err = service.UserRepository.FindByEmail(userRegisteration.Email)

	if err == nil {
		err = errors.New("Registered!!")
		return nil, err
	}

	user = &entities.User{}

	user.Email = userRegisteration.Email

	hash_utility := utilities.Hash{}
	if user.Password, err = hash_utility.Make(userRegisteration.Password); err != nil {
		return nil, err
	}

	if user, err = service.UserRepository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
