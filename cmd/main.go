package main

import (
	"fmt"
	"iban-validator/pkg/domain"
)

func main() {
	iban := "CV05766894889188535145823"
	if valid := domain.Validate(iban); valid {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
