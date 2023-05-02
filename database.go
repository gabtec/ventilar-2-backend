package main

import (
	"os"
	"strings"

	"github.com/withmandala/go-log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DbInstance is an instance of gorm db entity, export instance  ---> TO EXPORT MUST DEFINE OUTSIDE any func
var DbInstance *gorm.DB

// DatabaseConnect - will connect to database
func DatabaseConnect() {
	l := log.New(os.Stderr).WithColor()
	l.Info("Connecting to database...")
	
	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_SECRET") + " port=" + os.Getenv("DB_PORT") + " dbname=" + os.Getenv("DB_NAME")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		if strings.Contains(err.Error(), "3D000") {
			
			// log.Panic("You must create a database first")
			l.Fatal("You must create a database first")
		}
		// DB exists, but something went wrong
		l.Fatal("Error connecting to database! ")
	}
	

	// atention to the order of the relations: 1st Ward, 2nd User
	db.AutoMigrate(&Ward{})

	// OK
	l.Info("Connected to Postgres Database: \"" + os.Getenv("DB_NAME") + "\"")

	DbInstance = db
}