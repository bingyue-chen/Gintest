package services

import (
	"Gintest/src/repositories"
)

func NewAuthenticationService(userRepository repositories.UserRepositoryContract) AuthenticationService {

	return AuthenticationService{UserRepository: userRepository}
}

func NewRegistrationService(userRepository repositories.UserRepositoryContract) RegistrationService {

	return RegistrationService{UserRepository: userRepository}
}

func NewProfileService(userRepository repositories.UserRepositoryContract) ProfileService {

	return ProfileService{UserRepository: userRepository}
}
