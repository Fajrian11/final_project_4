package main

import (
	"final_project_4/router"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := router.StartAPP()
	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "9000" // Default port if not specified
	}
	err := r.Run()
	if err != nil {
		log.Println("[ERROR GRACEFUL]", err)
		os.Exit(1)
	}
}
