package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "host=db user=user password=password dbname=gotodo port=5432 sslmode=disable"

	var db *gorm.DB
	var err error

	// Retry connection for up to 30 seconds
	maxRetries := 30
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("DB connection is successful!")
			return db
		}

		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(1 * time.Second)
	}

	log.Fatal("Could not connect to database after all retries: ", err)
	return nil
}
