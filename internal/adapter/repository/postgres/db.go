package repository

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")
	DBPORT := os.Getenv("DB_PORT")

	// Connection string
	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	HOST, DBPORT, USER, PASS, DBNAME)

	fmt.Println("Connecting to database...")
	// Open a connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", HOST, USER, PASS, DBNAME, DBPORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	DB = db
	fmt.Println("Successfully connected!")
}
