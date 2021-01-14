# Somma

Scrivere un programma che legga da **riga di comando** una sequenza di numeri interi e stampi a video il risultato della somma dei numeri letti, calcolato utilizzando una funzione ricorsiva.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione ricorsiva `Somma(sl []int) (somma int)` che riceve in input un valore `[]int` nel parametro `sl` e restituisce un valore `int` pari alla somma dei valori presenti in `sl`.

##### Esempio d'esecuzione:

```text
$ go run somma.go 1 4 2 7
Somma valori: 14

$ go run somma.go 10     
Somma valori: 10

$ go run somma.go 10 4 24 17
Somma valori: 55

$ go run somma.go 1 ciao 5  
Somma valori: 6
```