# Laboratorio 7 - Array e slice I
## 1 Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	var a [4]int = [4]int{1, 2, 3, 4}

	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], " ")
	}

}
```
## 2 Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	var a [4]int = [4]int{1, 2, 3, 4}

	for i := 1; i <= len(a); i++ {
		fmt.Print(a[i], " ")
	}

}
```
## 3 Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	b := [3]rune{'a', 'b', 'c'}

	for _, v := range b {
		fmt.Printf("%c ", v)
	}

}
```
## 4 Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	var a [5]string

	a[1] = "hello"
	a[4] = "world"

	for i := range a {
		fmt.Print(a[len(a)-1-i])
	}

}
```
## 5 Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	var a [6]int

	for i := range a {
		a[i] = i
	}

	for _, v := range a {
		v *= 2
	}

	fmt.Println(a)
}
```
## 6 Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	var a [6]int

	for i := range a {
		a[i] = i
	}

	var b [6]int
	b = a

	for i := range b {
		b[i] = i * 2
	}

	fmt.Println(a)
}
```
## 7 Trova l'errore

Il seguente programma non può essere eseguito a causa di un errore. Qual è l'errore? Come è possibile risolvere il problema?

```go
package main

func main() {
	var n int = 4

	var a [n]int

	for i := 0; i < n; i++ {
		a[i] = i
	}
}
```
## 8 Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {

	var n int = 5

	var s []int
	s = make([]int, n)

	for i := 0; i < n; i++ {
		s[i] = i
	}

	fmt.Println(s)
}
```
## 9 Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {

	var s string = "Ma il cielo è sempre più blu!"

	var s2 []rune = []rune(s)

	fmt.Println(len(s), len(s2))
}
```
## 10 Stampa in ordine inverso


Scrivere un programma che, dopo aver letto da **standard input** un numero intero `n`, chiede all'utente di inserire `n` numeri interi (sempre da **standard input**). 

Il programma deve stampare gli `n` numeri interi in ordine inverso rispetto a quello di inserimento.

*Suggerimento:* Per creare dinamicamente una slice si utilizzi la funzione `make()`.

##### Esempio d'esecuzione:

```text
$ go run stampa_rovescio.go 
9
Inserisci 9 numeri:
1 -12 3 -4 5 -6 7 -7 9
Numeri in ordine inverso:
9 -7 7 -6 5 -4 3 -12 1 

$ go run stampa_rovescio.go 
5
Inserisci 5 numeri:
1 2 3 4 5
Numeri in ordine inverso:
5 4 3 2 1 
```
## 11 Gioco dei dadi

Scrivere un programma che simuli il gioco dei dadi tra un numero variabile di giocatori.
Il programma deve leggere da **standard input** il numero di giocatori, i loro nomi, e il numero di turni da giocare.
Il nome di ogni giocatore è una stringa senza spazi.

Ad ogni turno di gioco, il programma deve simulare due lanci di dado per ogni giocatore (usando la funzione `rand.Intn()` del package `math/rand`). Ogni turno è vinto dal giocatore la cui somma dei lanci è maggiore. In caso di parità vince il giocatore che ha giocato prima. L'ordine di gioco è uguale all'ordine di inserimento dei nomi dei giocatori.

Prima di terminare, il programma deve stampare un riepilogo con il vincitore di ogni turno.

*Suggerimento:* per inizializzare il generatore di numeri randomici e avere un risultato diverso ad ogni esecuzione, inserite l'istruzione
```go
rand.Seed(int64(time.Now().Nanosecond()))
```
all'interno della funzione `main()`.


```text
$ go run dadi.go
Inserisci il numero di giocatori: 4
Inserisci i nomi dei 4 giocatori:
Mario
Luca
Sofia
Marta
Inserisci il numero di turni: 3
=== Turno 1 ===
        Giocatore Mario, lanci di valore: 3 e 5
        Giocatore Luca, lanci di valore: 2 e 5
        Giocatore Sofia, lanci di valore: 3 e 3
        Giocatore Marta, lanci di valore: 1 e 6
        Turno 1, vincitore: Mario
=== Turno 2 ===
        Giocatore Mario, lanci di valore: 5 e 2
        Giocatore Luca, lanci di valore: 2 e 6
        Giocatore Sofia, lanci di valore: 5 e 2
        Giocatore Marta, lanci di valore: 6 e 5
        Turno 2, vincitore: Marta
=== Turno 3 ===
        Giocatore Mario, lanci di valore: 3 e 2
        Giocatore Luca, lanci di valore: 1 e 2
        Giocatore Sofia, lanci di valore: 3 e 4
        Giocatore Marta, lanci di valore: 4 e 1
        Turno 3, vincitore: Sofia
Vittorie:
Vincitore turno 1: Mario
Vincitore turno 2: Marta
Vincitore turno 3: Sofia
```
