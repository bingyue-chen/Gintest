package repositories

import (
	"Gintest/src/managers"
)

func NewUserRepository(databaseManager managers.DatabaseManagerContract) UserRepository {

	return UserRepository{DatabaseManager: databaseManager}
}
