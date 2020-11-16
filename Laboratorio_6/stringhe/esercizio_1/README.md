# Triangolo

Scrivere un programma che legga da **standard input** un numero intero `n` e, come mostrato nell'**Esempio di esecuzione**, stampi a video un triangolo rettangolo con base e altezza di lunghezza `n` utilizzando il carattere `*` (asterisco).

Se `n` è negativo o nullo, anziché stampare il triangolo il programma deve stampare un messaggio d'errore.

**Nota:** Si utilizzi un solo ciclo `for` ed una variabile di tipo `string`.

##### Esempio d'esecuzione:

```text
$ go run triangolo.go
-2
Dimensione non sufficiente

$ go run triangolo.go
5
*
**
***
****
*****
``` 