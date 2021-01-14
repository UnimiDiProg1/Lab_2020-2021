package main

import "fmt"
import "os"
import "unicode"

func main() {

	valori := os.Args[1:]

	for posizione, parola := range valori {
		fmt.Print(TrasformaParola(parola, posizione), " ")
	}

	fmt.Println()
}

func TrasformaParola(parola string, posizione int) (trasformata string) {
	var upper bool
	if posizione%2 == 0 {
		upper = true
	} else {
		upper = false
	}

	for _, v := range parola {

		if upper {
			trasformata += string(unicode.ToUpper(v))
		} else {
			trasformata += string(unicode.ToLower(v))
		}

		upper = !upper
	}

	return
}
