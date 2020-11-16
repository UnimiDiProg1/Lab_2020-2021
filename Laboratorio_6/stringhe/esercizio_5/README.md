# Somma stringhe

Scrivere un programma che:
1. Legga da **standard input** una sequenza di stringhe senza spazi, una per linea.
2. Converta ciascuna stringa in numero intero. A tal fine è possibile utilizzare la funzione 'strconv.Atoi' del package 'strconv'.
3. La lettura termina quando viene letta una stringa che non sia convertibile in intero (ovvero che contenga caratteri diversi da numeri interi) e viene stampata la somma dei numeri inseriti.

##### Esempio d'esecuzione:

```text
$ go run somma_stringhe.go
inserisci un intero:
3
inserisci un intero:
4
inserisci un intero:
a
somma: 7
```
