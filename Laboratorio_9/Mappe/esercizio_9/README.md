# Sottosequenze (1)

Scrivere un programma che legga da **riga di comando** una sequenza di numeri interi e stampi tutte le sottosequenze che iniziano e finiscono con lo stesso numero. Ciascuna sottosequenza deve essere stampata su una riga diversa.

##### Esempio:
 
Se la sequenza di input è `1 2 3 -14 2 5`, l'unica sottosequenza è `2 3 -14 2`. Se la sequenza di input è `1 2 1 2 3`, abbiamo 2 sottosequenze: `1 2 1` e `2 1 2`.

Si consideri che:

  * se a **riga di comando** non viene specificata alcuna sequenza, il programma non deve stampare nulla;
  * una sottosequenza può essere contenuta in una sottosequenza più grande;
  * ogni sottosequenza deve comparire una sola volta tra quelle stampate a video;
  * le sottosequenze devono essere stampate in ordine di lunghezza (dalla più corta alla più lunga).
  
##### Esempio d'esecuzione:

```text
$ go run sottosequenze.go 1 2 3 4 2 5
2 3 4 2

$ go run sottosequenze.go 1 2 1 2 3
1 2 1
2 1 2

$ go run sottosequenze.go 1 2 -45 2 1
2 -45 2
1 2 -45 2 1

$ go run sottosequenze.go 

$ go run sottosequenze.go 10 3 12 10 4 5 13
10 3 12 10

$ go run sottosequenze.go 1 2 5 5 2 3 1
5 5
2 5 5 2
1 2 5 5 2 3 1

$ go run sottosequenze.go 1 3 2 2 3 2 2 3 1
2 2
2 3 2
2 3 2 2
3 2 2 3
2 2 3 2
2 2 3 2 2
3 2 2 3 2 2 3
1 3 2 2 3 2 2 3 1
```