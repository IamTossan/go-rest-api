package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	minPort = 1024
	maxPort = 65535
)

func main() {
	port := flag.Uint("p", 8080, "the port used by this app")
	flag.Parse()

	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_DBNAME"),
	)

	if *port <= minPort || *port > maxPort {
		log.Fatalf("incorrect port : %d", *port)
	}

	fmt.Printf("App listening to port : %d\n", *port)

	p := fmt.Sprintf(":%d", *port)
	a.Run(p)
}
