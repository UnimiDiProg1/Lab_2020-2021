# Numeri pari superiori al minimo dispari

Scrivere un programma che legga da **riga di comando** una sequenza di valori positivi. Sia `minDispari` il minimo valore dispari tra i valori interi letti.
Il programma deve stampare a video i valori interi pari, e superiori a `minDispari`, presenti nella sequenza di valori letta.

Si assuma che tra i valori interi letti da riga di comando sia sempre presente almeno un valore dispari.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `MinimoDispari(sl []int) int` che riceve in input un valore `[]int` nel parametro `sl` e restituisce il minimo valore dispari presente in `sl`.

##### Esempio d'esecuzione:

```text
$ go run pari.go 3 2 4      
4 

$ go run pari.go 2 5 4 6 7 8
6 8 

$ go run pari.go 2 6 8 1    
2 6 8 

$ go run pari.go 1 3 5  

```