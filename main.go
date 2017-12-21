package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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

	if *port < 1000 || *port > 9999 {
		log.Fatalf("incorrect port : %d", *port)
	}

	fmt.Printf("App listening to port : %d\n", *port)

	p := ":" + fmt.Sprint(*port)
	a.Run(p)
}
