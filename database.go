package main

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	Host string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

func NewDBConfig(dbHost, user, password, dbName, port, sslMode string) *DBConfig {
	return &DBConfig{
		Host: dbHost,
		User:     user,
		Password: password,
		DBName:   dbName,
		Port:     port,
		SSLMode:  sslMode,
	}
}

func (dc DBConfig) Connect() (*gorm.DB, error) {
	dsn := dc.createDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (dc DBConfig) createDSN() string {
	return "host=" + dc.Host + " user=" + dc.User + " password=" + dc.Password + " dbname=" + dc.DBName + " port=" + dc.Port + " sslmode=" + dc.SSLMode
}

func Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// 20240429_add_user_table.go
		AddUserTable(),
	})
	if err := m.Migrate(); err != nil {
		return err
	}

	return nil
}
