package postgresdb

import (
	"fmt"
	"log"
	"rest-api/config"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	var err error

	p := config.Config("DB_PORT")

	// because our config function returns a string, we are parsing our str to int here - hack ??
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"),
		port, config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(
		"Connect to Postgres Successfully! With: " +
			config.Config("DB_HOST") +
			":" +
			p +
			"/" +
			config.Config("DB_NAME"),
	)
	return nil
}
