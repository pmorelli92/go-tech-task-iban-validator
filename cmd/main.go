package main

import (
	"iban-validator/pkg/http"
	"os"
)

func main() {

	// If port is not assigned it will start on a random port
	port := os.Getenv("PORT")

	if err := http.Bootstrap(port); err != nil {
		panic(err)
	}
}
