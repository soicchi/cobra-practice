package main

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// Migrations is a list of migrations
func AddUserTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20240429_add_user_table",
		Migrate: func(tx *gorm.DB) error {
			type user struct {
				gorm.Model
				Name string `gorm:"size:255;not null;unique"`
			}

			return tx.Migrator().CreateTable(&user{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("users")
		},
	}
}
