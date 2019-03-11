package main

import (
	"fmt"
	"iban-validator/pkg/domain"
)

func main() {
	iban := "CV05766894889188535145823"
	if err := domain.Validate(iban); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("OK")
	}
}
