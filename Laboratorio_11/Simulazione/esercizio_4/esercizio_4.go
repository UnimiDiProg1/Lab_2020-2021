package main

import "fmt"
import "strconv"
import "math"
import "os"

func main() {

	n := os.Args[1]
	d, _ := strconv.Atoi(os.Args[2])

	minimo := math.MaxInt64

	for contatore := 0; contatore < int(math.Pow(2, float64(len(n)))); contatore++ {
		combinazione := fmt.Sprintf("%0"+strconv.Itoa(len(n))+"b", contatore)

		var numero string
		for indice, cifraBinaria := range combinazione {
			if cifraBinaria == '1' {
				numero += string(n[indice])
			}
		}

		if len(numero) == len(n)-d {
			numeroConvertito, _ := strconv.Atoi(numero)
			if numeroConvertito < minimo {
				minimo = numeroConvertito
			}
		}
	}

	fmt.Println("numero migliore:", minimo)

}
