package domain

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

func Validate(iban string) error {
	charValues := *getCharValues()
	countries := *getCountries()

	ibanRune := []rune(iban)
	country := string(ibanRune[0:2])

	if length, ok := countries[country]; ok {
		if len(iban) == length {
			concat := ""

			// get the first 4 characters and put it at the end
			ibanRune = append(ibanRune, ibanRune[0:4]...)

			// now take from the 4 position onwards
			for _, rn := range ibanRune[4:] {
				current := string(rn)

				// if it is not a number then we take the string conversion
				if _, err := strconv.Atoi(current); err == nil {
					concat += current
				} else {
					concat += charValues[current]
				}
			}

			n := new(big.Int)
			n.SetString(concat, 10)
			res := n.Mod(n, big.NewInt(97))

			// concat % 97 == 1
			if res.String() == "1" {
				return nil
			}
		}
	}

	errMsg := fmt.Sprintf("The IBAN %s is not valid", iban)
	return errors.New(errMsg)
}
