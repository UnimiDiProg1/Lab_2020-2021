package main

import (
	"fmt"
)

func main() {

	var numero int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	// stampa n righe di caratteri ** e ++ alternati sia sulla stessa riga che sulla colonna
	var i, j int
	for i = 0; i < numero; i++ {
		for j = 0; j < numero; j++ {
			// Un modo per risolvere questo problema è quello di considerare 4 casi distinti:
			// caso 1 e 2 -> stampa di un *
			// caso 3 e 4 -> stampa di un +
			// I casi sono 4 perché il pattern che si ripete è di 4 caratteri (**++)

			// Un modo per farlo è quello di distinguere le righe di indice pari e dispari e fare una stampa
			// differenziata
			// Il modo proposto in questo codice è invece un'altro:
			// dato che quello che cambia tra una riga e l'altra è che i simboli **++ sono spostati di due posizioni,
			// mantengo una nuova variabile chiamata offset che sarà identica a j se siamo in una riga di indice pari,
			// mentre sarà j+2 (quindi spostata di due posizioni) se l'indice di riga è dispari
			var offset int = j
			if i%2 != 0 {
				offset += 2
			}

			if offset%4 <= 1 {
				fmt.Print("* ")
			} else {
				fmt.Print("+ ")
			}
		}
		fmt.Println()
	}

}
