package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConn() *gorm.DB {
	db, err := gorm.Open(
		postgres.Open(
			"host=localhost user=postgres password=postgres dbname=go_grpc port=5432 sslmode=disable",
		),
		&gorm.Config{},
	)
	if err != nil {
		panic(err)
	}
	return db
}
