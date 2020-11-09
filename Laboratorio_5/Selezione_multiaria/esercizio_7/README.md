# Soluzione di equazioni di secondo grado

Scrivere un programma che legga da **standard input** tre numeri reali, `a`, `b` e `c`, che rappresentano i coefficienti di un'equazione di secondo grado (con incognita `x`) espressa nella forma `ax² + bx + c = 0`.
Il programma deve calcolare e stampare a video il valore delle radici dell'equazione.

*Suggerimento:* Nel caso in cui l'equazione non ammetta soluzioni reali, potete memorizzare i valori delle radici in variabili di tipo `complex`. 

**Nota:** Si utilizzi il costrutto `switch case`.

##### Esempio d'esecuzione:

```text
$ go run equazioni.go 
1 5 4
Due radici reali -1 -4

$ go run equazioni.go
1 0 4
Due radici complesse (0+2i) (-0-2i)

$ go run equazioni.go
1 4 4
Due radici reali coincidenti -2 -2
``` 