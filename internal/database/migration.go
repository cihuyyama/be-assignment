package database

import (
	"be-assignment/domain"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&domain.User{},
		&domain.PaymentAccount{},
		&domain.PaymentHistory{},
	)
}
