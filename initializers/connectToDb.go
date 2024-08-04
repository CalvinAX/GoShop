package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	fmt.Println("INFO ##### Connecting with database...")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, databaseName, port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("ERROR ##### Couldn't connect to database! Is it running?")
		panic(err)
	} else {
		fmt.Println("INFO ##### Successfully connected to the database!")
	}
}
