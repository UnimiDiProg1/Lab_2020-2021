# Soluzione di equazioni di primo grado

Scrivere un programma che legga da **standard input** due numeri reali, `a` e `b`, che rappresentano i coefficienti di un'equazione
di primo grado (con incognita `x`) espressa nella forma `ax + b = 0`.
Il programma deve calcolare e stampare a video il valore della radice dell'equazione (il valore di `x` per cui l’uguaglianza risulta verificata). 

*Suggerimento:* Il valore della radice è pari a `-b/a`.

##### Esempio d'esecuzione:

```text
$ go run equazione.go
3 6
La radice è -2

$ go run equazione.go
4 2
La radice è -0.5

$ go run equazione.go
10 -4
La radice è 0.4

$ go run equazione.go
10 -10
La radice è 1
``` 

Cosa succede se `a=0` e `b≠0`?

Cosa succede se `a=0` e `b=0`?