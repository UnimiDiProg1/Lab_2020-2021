package main

import (
	"fmt"
	"strings"
)

func main() {

	for {
		var stringaInput string
		fmt.Println("inserisci una stringa:")
		fmt.Scanln(&stringaInput)

		if stringaInput == "" {
			break
		}

		fmt.Println("stringa in maiuscolo:",strings.ToUpper(stringaInput))
	}

	fmt.Println("ciao")
}
