package main

import (
	"os"
	"strings"

	gou "github.com/gabtec/gabtec-gou"
	"github.com/gabtec/ventilar-2-backend/types"
	"github.com/withmandala/go-log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DbInstance is an instance of gorm db entity, export instance  ---> TO EXPORT MUST DEFINE OUTSIDE any func.
// var DbInstance *gorm.DB

// DatabaseConnect - will connect to database.
func DatabaseConnect() *gorm.DB {
	l := log.New(os.Stderr).WithColor()
	l.Info("Connecting to database...")

	// dsn := "host=localhost user=admin password=admin dbname=atlas_db port=5432 sslmode=disable TimeZone=Europe/Lisbon"
	// OR
	// dsn := "postgres://admin:admin@localhost:5432/atlas_db"
	url := gou.GetEnv("DB_URL", "")

	dsn := TrimNewLine(url)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		if strings.Contains(err.Error(), "3D000") {
			l.Fatal("You must create a database first")
		}
		// DB exists, but something went wrong
		l.Fatal("Error connecting to database! ")
	}

	if readArguments(os.Args) != "production" {
		// atention to the order of the relations: 1st Ward, 2nd User
		err = db.AutoMigrate(&types.Ward{}, &types.User{})

		if err != nil {
			l.Fatal("Database automigration of models, failed.")
		}
	}

	// OK
	var dbName string
	db.Raw("SELECT current_database() as dbName;").Row().Scan(&dbName)
	l.Info("Connected  to \"" + dbName + "\" (postgresql)")

	return db
}

// TrimNewLine - ajshajsd
func TrimNewLine(s string) string {
	suffix := "\n"
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
