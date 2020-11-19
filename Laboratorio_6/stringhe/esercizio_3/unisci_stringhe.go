package main

import "fmt"

func main() {

	var stringaCompleta string

	for {
		var stringaInput string
		fmt.Println("inserisci una stringa:")
		fmt.Scanln(&stringaInput)

		if stringaInput == "" {
			break
		}
		
		stringaCompleta+= " " + stringaInput
	}

	fmt.Println("stringa completa:" + stringaCompleta)
}
