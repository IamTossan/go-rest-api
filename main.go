package main

import (
	"os"
)

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_DBNAME"),
	)

	a.Run(":8080")
}
