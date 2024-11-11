package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// TODO: Move to more reliable migration tool
	DB.AutoMigrate(&Book{}, &User{})
}

func CloseDB() {
	conn, err := DB.DB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	conn.Close()
}
