# Controlla sequenza (1)

Scrivere un programma che legga da **standard input** una stringa di caratteri in cui i caratteri corrispondenti a cifre decimali rappresentano dei numeri interi minori di 10.

All'interno della stringa letta, i caratteri corrispondenti a cifre decimali sono intervallati tra loro da sottostringhe formate da caratteri arbitrari.

Il programma deve considerare la sequenza di numeri interi nascosta all'interno della stringa letta da **standard input**.

**Esempio:**

All'interno della stringa di caratteri

`&4&$4%mamma5!6mia6cosa1succede0`

la sequenza di numeri interi nascosta è:

`4 4 5 6 6 1 0`

Il programma deve controllare che i numeri presenti all'interno della sequenza siano ordinati in senso decrescente, cioè dal più grande al più piccolo.

Nel caso in cui i numeri interi siano ordinati in senso decrescente, il programma deve stampare a video il messaggio `Sequenza nascosta ordinata.`.
In caso contrario, il programma deve stampare a video il messaggio `Sequenza nascosta non ordinata.`.

Si assuma che:
* all'interno della stringa di caratteri letta da **standard input** non siano presenti caratteri consecutivi corrispondenti a cifre decimali;
* all'interno della stringa di caratteri letta da **standard input** non sia presente nessun carattere di spaziatura, ossia un carattere il cui codice Unicode, passato come argomento alla funzione `func IsSpace(r rune) bool` del package `unicode`, fa restituire `true` alla funzione.

##### Esempio d'esecuzione:
```text
$ go run controlla_sequenza.go
abc8dèf6asa2sd0
Sequenza nascosta ordinata.

$ go run controlla_sequenza.go
Èbc8d8e5a#4§d1e
Sequenza nascosta ordinata.

$ go run controlla_sequenza.go
a°c9d8e6aà7sd1e
Sequenza nascosta non ordinata.
```
