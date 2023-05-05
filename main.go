package main

import (
	"os"

	gou "github.com/gabtec/gabtec-gou"
	"github.com/joho/godotenv"
	"github.com/withmandala/go-log"
	"gorm.io/gorm"
)

/*
* ============================
*    			MAIN block
* ============================.
 */

// DbInstance - the database conn instance, exported
var DbInstance *gorm.DB

func main() {
	l := log.New(os.Stderr).WithColor()

	// when running this app one argument (--env, -e)
	// is accepted
	// e.g. go run . -e [ production | development ]
	// it defaults to "development" if no arg provided

	if readArguments(os.Args) != "production" {
		// in production will use exported env variables like in docker or kubernetes
		// in development will use dotenv file
		err := godotenv.Load()
		if err != nil {
			l.Fatal("Error loading .env file")
		}
	}

	PORT := ":" + gou.GetEnv("API_PORT", "4000")

	DbInstance = DatabaseConnect()

	app := SetupFiber()
	l.Info("API url: http://localhost" + PORT + "/api/")

	// Start Server
	l.Fatal(app.Listen(PORT))
}

/*
* ============================
*    			Helper block
* ============================
 */

// ReadArguments - read command line --env arg, or defaults to "development"
func readArguments(a []string) string {
	runningEnv := "development" // default
	envIndex := 0

	for idx, arg := range a {
		if arg == "-e" || arg == "--env" {
			envIndex = idx + 1
		}
	}

	if envIndex < 2 {
		return runningEnv
		// return "", errors.New("--env argument not found")
	}

	runningEnv = a[envIndex]

	return runningEnv
}
