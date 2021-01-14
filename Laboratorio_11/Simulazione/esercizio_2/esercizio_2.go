package main

import "fmt"
import "os"

func main() {

	input := os.Args[1] // abba

	sottostringhe := make(map[string]int)

	for inizio := 0; inizio < len(input); inizio++ {
		for fine := inizio + 1; fine < len(input); fine++ {

			if input[inizio] == input[fine] {
				sottostringa := input[inizio : fine+1]
				sottostringhe[sottostringa]++
			}

		}

	}

	for lunghezza := len(input); lunghezza >= 3; lunghezza-- {
		for sottostringa, contatore := range sottostringhe {
			if len(sottostringa) == lunghezza {
				fmt.Printf("%s -> Occorrenze: %d\n", sottostringa, contatore)
			}
		}
	}
}
