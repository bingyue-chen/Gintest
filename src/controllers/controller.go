package controllers

import (
	"Gintest/src/managers"
	"Gintest/src/repositories"
	"Gintest/src/services"
)

type Provider struct {
	instances map[string]interface{}
}

var provider *Provider

func GetProvider() *Provider {

	if provider == nil {
		provider = &Provider{instances: map[string]interface{}{}}
	}

	return provider
}

func GetNewProvider() *Provider {
	provider = &Provider{instances: map[string]interface{}{}}

	return provider
}

func (p *Provider) GetAuthenticationService() services.AuthenticationServiceContract {

	if instance, ok := p.instances["AuthenticationService"]; ok {
		return instance.(services.AuthenticationServiceContract)
	}

	authenticationService := services.NewAuthenticationService(p.GetUserRepository())

	p.SetInstance("AuthenticationService", &authenticationService)

	return &authenticationService
}

func (p *Provider) GetRegistrationService() services.RegistrationServiceContract {

	if instance, ok := p.instances["RegistrationService"]; ok {
		return instance.(services.RegistrationServiceContract)
	}

	registrationService := services.NewRegistrationService(p.GetUserRepository())

	p.SetInstance("RegistrationService", &registrationService)

	return &registrationService
}

func (p *Provider) GetProfileService() services.ProfileServiceContract {

	if instance, ok := p.instances["ProfileService"]; ok {
		return instance.(services.ProfileServiceContract)
	}

	profileService := services.NewProfileService(p.GetUserRepository())

	p.SetInstance("ProfileService", &profileService)

	return &profileService
}

func (p *Provider) GetUserRepository() repositories.UserRepositoryContract {

	if instance, ok := p.instances["UserRepository"]; ok {
		return instance.(repositories.UserRepositoryContract)
	}

	userRepository := repositories.NewUserRepository(p.GetDatabaseManager())

	p.SetInstance("UserRepository", &userRepository)

	return &userRepository
}

func (p *Provider) GetDatabaseManager() managers.DatabaseManagerContract {

	if instance, ok := p.instances["DatabaseManager"]; ok {
		return instance.(managers.DatabaseManagerContract)
	}

	databaseManager := managers.NewDatabaseManager()

	p.SetInstance("DatabaseManager", &databaseManager)

	return &databaseManager
}

func (p *Provider) SetInstance(name string, instance interface{}) {
	p.instances[name] = instance
}
