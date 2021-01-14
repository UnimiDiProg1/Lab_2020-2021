package main

import "fmt"
import "os"
import "strconv"
import "strings"

func main() {

	valori := os.Args[1:]

	for indice := 0; indice < len(valori); indice += 2 {

		n, _ := strconv.Atoi(valori[indice+1])

		//for ripetizioni := 0; ripetizioni < n; ripetizioni ++ {
		//	fmt.Print(valori[indice])
		//}
		fmt.Print(strings.Repeat(valori[indice], n))

	}

	fmt.Println()

}
