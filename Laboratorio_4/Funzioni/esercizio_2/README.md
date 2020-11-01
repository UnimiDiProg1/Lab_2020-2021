# Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import "fmt"

func test(x int) (y int) {
	/* Le variabili 'a' e 'b' definite nel 'main' sono out of scope */
	/* La seguente istruzione genererebbe un errore di compilazione */
	//fmt.Println(a, b)
	
	fmt.Println("Funzione 'test' - Inizio.")
	y = x * 10
	fmt.Println("Funzione 'test' - Prima dell'ultima istruzione.")
	return
	//fmt.Println("Funzione 'test' - Dopo dell'ultima istruzione.")
}

func main() {
	fmt.Println("Funzione 'main' - Inizio.")
	var a, b int = 10, 5
	fmt.Println("Funzione 'main' - Valori di 'a' e 'b' prima di chiamare/invocare " +
		"la funzione 'test'.")
	fmt.Println(a, b)
	b = test(a)
	fmt.Println("Funzione 'main' - Valori di 'a' e 'b' dopo aver chiamato/invocato " +
		"la funzione 'test'.")
	fmt.Println(a, b)
	fmt.Println("Funzione 'main' - Fine.")	
}
```