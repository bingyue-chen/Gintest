package managers

import (
	"database/sql"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"Gintest/src/entities"
	"Gintest/src/utilities"
)

type DatabaseManagerContract interface {
	Connect()
	Getclient() *gorm.DB
}

type DatabaseManager struct {
	Client *gorm.DB
}

func (manager *DatabaseManager) Connect() {

	host := utilities.Getenv("DB_HOST")
	port := utilities.Getenv("DB_PORT")
	database := utilities.Getenv("DB_DATABASE")
	username := utilities.Getenv("DB_USERNAME")
	pasword := utilities.Getenv("DB_PASSWORD")

	dsn := "user=" + username + " password=" + pasword + " dbname=" + database + " host=" + host + " port=" + port + " sslmode=disable TimeZone=UTC"

	DatabaseClient, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	dbPool, err := DatabaseClient.DB()

	if err != nil {
		panic("Failed retrive db pool on Client!")
	}

	configConnectionPool(dbPool)

	//DatabaseClient.AutoMigrate(&entities.User{})

	migrate(DatabaseClient)

	manager.Client = DatabaseClient
}

func (manager *DatabaseManager) Getclient() *gorm.DB {
	return manager.Client
}

func configConnectionPool(dbPool *sql.DB) {

	dbPool.SetMaxIdleConns(10)

	dbPool.SetMaxOpenConns(100)

	dbPool.SetConnMaxLifetime(time.Hour)
}

func migrate(dbClient *gorm.DB) {

	dbClient.AutoMigrate(&entities.User{})
}
