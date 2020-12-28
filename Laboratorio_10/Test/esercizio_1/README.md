# Pari e dispari

Scrivere un programma che legga da **standard input** un numero intero `n` e stampi a video se `n` è pari o dispari.

Oltre alla funzione `main()`, devono essere definite le seguenti funzioni:
* una funzione `Pari(n int) bool` che riceve in input un valore intero `n` e restituisce `true` se `n` è pari, `false` altrimenti;
* una funzione `Dispari(n int) bool` che riceve in input un valore intero `n` e restituisce `true` se `n` è dispari, `false` altrimenti.

Verificare il corretto funzionamento delle funzioni utilizzando le unit di test allegate.

##### Esempio d'esecuzione:

```text
$ go run pari_dispari.go
Inserisci un numero: 5       
Il numero 5 è dispari

$ go run pari_dispari.go
Inserisci un numero: 4
Il numero 4 è pari
```