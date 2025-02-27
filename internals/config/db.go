package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	URL        string
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
	DbPort     string
	DbSslMode  string
}

var config = Config{}

func Connect() (*gorm.DB, error) {
	config.Read()

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.DbUser,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbName,
		config.DbSslMode,
	)

	conn, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB = conn

	return conn, err

}

func (c *Config) Read() {
	config.DbHost = os.Getenv("DB_HOST")
	config.DbUser = os.Getenv("DB_USER")
	config.DbPassword = os.Getenv("DB_PASSWORD")
	config.DbName = os.Getenv("DB_NAME")
	config.DbPort = os.Getenv("DB_PORT")
	config.DbSslMode = os.Getenv("DB_SSLMODE")
}
