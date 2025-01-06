// By Saif Hamdan, Team Lead
// Date: 2025/1/6
package data

import (
	models "github.com/saifhamdan/go-apigateway-blueprint/models/v1"
	"github.com/saifhamdan/go-apigateway-blueprint/pkg/db"
	"github.com/saifhamdan/go-apigateway-blueprint/pkg/logger"
	"gorm.io/gorm"
)

func Seed(logger *logger.Logger, db *db.DB) {

	tx := db.Postgres.DB.Begin()

	err := seedUsers(tx)
	if err != nil {
		tx.Rollback()
		logger.Fatal(err)
	}

	tx.Commit()
}

func seedUsers(tx *gorm.DB) error {
	// create users
	users := []*models.User{
		{
			Username:     "admin",
			UserPassword: "admin", // password should be hashed and salted but for the sake of simplicity we will use plain text
			FirstName:    "Admin",
			FamilyName:   "",
			Email:        "admin@company.com",
			Phone:        "123456789",
		},
		{
			Username:     "user",
			UserPassword: "user", // password should be hashed and salted but for the sake of simplicity we will use plain text
			FirstName:    "User",
			FamilyName:   "",
			Email:        "user@company.com",
			Phone:        "123456789",
		},
	}

	return tx.Save(users).Error
}
