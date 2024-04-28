package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	dbConfig := NewDBConfig(os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), "disable")
	db, err := dbConfig.Connect()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connected to database")

	if err := Migrate(db); err != nil {
		log.Fatal(err)
	}
	log.Printf("Migrated database")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
