package services

import (
	"Gintest/src/entities"
	"Gintest/src/repositories"
	"Gintest/src/utilities"
)

type ProfileServiceContract interface {
	Retrive(id int64) (*entities.User, error)
	Update(user *entities.User, data *entities.UserUpdation) (*entities.User, error)
}

type ProfileService struct {
	UserRepository repositories.UserRepositoryContract
}

func (service *ProfileService) Retrive(id int64) (*entities.User, error) {
	return service.UserRepository.Find(id)
}

func (service *ProfileService) Update(user *entities.User, userUpdation *entities.UserUpdation) (*entities.User, error) {

	updatedData := map[string]interface{}{}

	if userUpdation.FirstName != "" {
		updatedData["first_name"] = userUpdation.FirstName
	}

	if userUpdation.LastName != "" {
		updatedData["last_name"] = userUpdation.LastName
	}

	if userUpdation.Password != "" {

		hash_utility := utilities.Hash{}
		password, err := hash_utility.Make(userUpdation.Password)

		if err != nil {
			return nil, err
		}

		updatedData["password"] = password
	}

	return service.UserRepository.Update(user, updatedData)
}
