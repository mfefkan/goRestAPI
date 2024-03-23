package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	return gorm.Open(postgres.Open("postgres://root:root@localhost:5432/test"), &gorm.Config{})
}

// connection string =>   "docker run -e POSTGRES_PASSWORD="root" -e POSTGRES_USER="root" -e POSTGRES_DB="test" -p 5432:5432 postgres"
