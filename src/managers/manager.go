package managers

func NewDatabaseManager() DatabaseManager {

	databaseManager := DatabaseManager{}

	databaseManager.Connect()

	return databaseManager
}
