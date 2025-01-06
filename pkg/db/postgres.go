// By Saif Hamdan, Team Lead
// Date: 2025/1/6
package db

import (
	"fmt"

	models "github.com/saifhamdan/go-apigateway-blueprint/models/v1"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type postgres struct {
	*gorm.DB
}

func (d *DB) newPostgres() (*postgres, error) {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		d.cfg.PostgresHost,
		d.cfg.PostgresUser,
		d.cfg.PostgresPassword,
		d.cfg.PostgresName,
		d.cfg.PostgresPort,
	)

	d.logger.Infof("connecting to postgres: %s", dns)

	db, err := gorm.Open(postgresDriver.Open(dns), &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "blueprint_",
		},
	})
	if err != nil {
		return nil, err
	}

	d.logger.Infof("postgres connection established %s", dns)

	return &postgres{DB: db}, err
}

func (p *postgres) Close() error {
	sqlDB, err := p.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (db *postgres) MigrateDB() error {
	err := db.AutoMigrate(
		// common models
		models.User{},
	)

	return err
}
