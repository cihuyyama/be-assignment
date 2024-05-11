package database

import (
	"be-assignment/internal/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabaseConnection(cnf *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s"+
			" port=%s"+
			" user=%s"+
			" password=%s"+
			" dbname=%s"+
			" sslmode=disable",
		cnf.DB.Host,
		cnf.DB.Port,
		cnf.DB.User,
		cnf.DB.Password,
		cnf.DB.Name,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	println("Connected to database")

	return db
}
