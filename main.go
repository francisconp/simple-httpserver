package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// Define command line paramiters
	port := flag.String("port", "", "Define default port")
	dir := flag.String("dir", "", "Define default directory")
	flag.Parse()

	// Define work directory
	if err := os.Chdir(*dir); err != nil {
		log.Fatalf("Command line example: simple-httpserver -port 8080 -dir /directory/  %s: %v", *dir, err)
	}

	// Define a handler
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	log.Printf("Directory: %s, Port: %s\n", *dir, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
