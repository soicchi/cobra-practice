package main

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

func NewDBConfig(user, password, dbName, port, sslMode string) *DBConfig {
	return &DBConfig{
		User:     user,
		Password: password,
		DBName:   dbName,
		Port:     port,
		SSLMode:  sslMode,
	}
}

func (dc DBConfig) Connect() (*gorm.DB, error) {
	dsn := dc.createDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (dc DBConfig) createDSN() string {
	return "host=localhost user=" + dc.User + " password=" + dc.Password + " dbname=" + dc.DBName + " port=" + dc.Port + " sslmode=" + dc.SSLMode
}

func (dc DBConfig) Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{})
	if err := m.Migrate(); err != nil {
		return err
	}

	return nil
}
