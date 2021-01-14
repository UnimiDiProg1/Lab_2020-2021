# Prodotto scalare

Scrivere un programma che:
1. Legga da **standard input** un numero intero `n`.
2. Chieda all'utente di inserire, sempre da **standard input**, due sequenze di `n` numeri interi:

    a<sub>0</sub> a<sub>1</sub> a<sub>2</sub> ...  a<sub>n</sub>

    b<sub>0</sub> b<sub>1</sub> b<sub>2</sub> ...  b<sub>n</sub>

3. Stampi a video il risultato dell'espressione

    a<sub>0</sub> * b<sub>0</sub> + a<sub>1</sub> * b<sub>1</sub> + ... + a<sub>n</sub> * b<sub>n</sub>

    calcolato utilizzando una funzione ricorsiva.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione ricorsiva `Moltiplica(sl1, sl2 []int) int` che riceve in input due valori `[]int`, rispettivamente nei parametri `sl1` e `sl2`, e restituisce un valore `int` pari a `sl1[0]*sl2[0] + sl1[1]*sl2[1] + ... + sl1[n-1]*sl2[n-1]`, dove `n` è pari a `len(sl1)` e `len(sl2)`.  

##### Esempio d'esecuzione:
```text
$ go run prodotto_scalare.go 
4
1 4 2 7
2 3 8 9
Prodotto = 93
``` 
  