package main

import "iban-validator/pkg/http"

func main() {
	if err := http.Bootstrap(); err != nil {
		panic(err)
	}
}