# Rombo

Scrivere un programma che legga da **standard input** un numero intero `n` e, come mostrato nell'**Esempio di esecuzione**, stampi a video un rombo con diagonale maggiore e diagonale minore di lunghezza `2*n+1` utilizzando il carattere `*` (asterisco).

A tal fine definire due stringhe: 'stringaSpazi' e 'stringaAsterischi' contenenti il massimo numero dei caratteri necessari e utilizzare le loro sottostringhe per la stampa.

Se `n` è negativo o nullo, anziché stampare il rombo il programma deve stampare un messaggio d'errore.

##### Esempio d'esecuzione:

```text
$ go run rombo.go
3
   *
  ***
 *****
*******
 *****
  ***
   *

$ go run rombo.go 
0
Dimensione non sufficiente
``` 
