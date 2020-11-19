package main

import (
	"fmt"
)

func main() {

	var ultimo string

	for {
		var stringaInput string
		fmt.Println("inserisci un cognome:")
		fmt.Scanln(&stringaInput)

		if stringaInput == "" {
			break
		}
		
		if stringaInput > ultimo {
			ultimo = stringaInput
		}
	}

	fmt.Println("ultimo cognome:", ultimo)
}
