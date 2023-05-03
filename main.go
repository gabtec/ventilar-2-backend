package main

import (
	"os"

	"github.com/withmandala/go-log"

	"github.com/joho/godotenv"
)

/*
* ============================
*    			MAIN block
* ============================
 */
func main() {
	l := log.New(os.Stderr).WithColor()
	err := godotenv.Load()

	if err != nil {
		l.Fatal("Error loading .env file")
	}

	PORT := ":" + os.Getenv("API_PORT")

	DatabaseConnect()

	app := SetupFiber()
	l.Info("API Port is: localhost" + PORT)

	// Start Server
	l.Fatal(app.Listen(PORT))
}

/*
* ============================
*    			MAIN block
* ============================
 */
