# Laboratorio 6 - Stringhe
## 1 Triangolo

Scrivere un programma che legga da **standard input** un numero intero `n` e, come mostrato nell'**Esempio di esecuzione**, stampi a video un triangolo rettangolo con base e altezza di lunghezza `n` utilizzando il carattere `*` (asterisco).

Se `n` è negativo o nullo, anziché stampare il triangolo il programma deve stampare un messaggio d'errore.

**Nota:** Si utilizzi un solo ciclo `for` ed una variabile di tipo `string`.

##### Esempio d'esecuzione:

```text
$ go run triangolo.go
-2
Dimensione non sufficiente

$ go run triangolo.go
5
*
**
***
****
*****
``` 
## 2 Ripetizione

Scrivere un programma che:
1. legga da **standard input** un numero intero `n` ed una stringa senza spazi;
2. stampi `n` volte la stringa letta, intervallando le `n` occorrenze della stringa con il carattere `-` (meno).

##### Esempio d'esecuzione:

```text
$ go run ripetizione.go
5 test
test-test-test-test-test
```

## 3 Unisci stringhe

Scrivere un programma che:
1. Legga da **standard input** una sequenza di stringhe senza spazi, una per linea.
2. La lettura termina quando viene letta una stringa vuota `""` e viene stampata una stringa che le unisce tutte (separate con spazi).

##### Esempio d'esecuzione:

```text
$ go run unisci_stringhe.go
inserisci una stringa:
hello
inserisci una stringa:
world!
inserisci una stringa:

stringa completa: hello world!
```

## 4 Maiuscole

Scrivere un programma che:
1. Legga da **standard input** una sequenza di stringhe senza spazi, una per linea.
2. Stampi ciascuna stringa convertita in maiuscolo. A tal fine è possibile utilizzare la funzione 'strings.ToUpper' del package 'strings'.
3. La lettura termina quando viene letta una stringa vuota `""` e viene stampato 'ciao'.

##### Esempio d'esecuzione:

```text
$ go run maiuscole.go
inserisci una stringa:
pippo
stringa in maiuscolo: PIPPO
inserisci una stringa:

ciao
```

## 5 Somma stringhe

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

## 6 Ultimo cognome

Scrivere un programma che:
1. Legga da **standard input** una sequenza di cognomi (stringhe) senza spazi, uno per linea.
2. La lettura termina quando viene letta una stringa vuota `""` e viene stampato l'ultimo cognome in ordine alfabetico.

##### Esempio d'esecuzione:

```text
$ go run ultimo_cognome.go
inserisci un cognome:
rossi
inserisci un cognome:
verdi
inserisci un cognome:
bianchi
inserisci un cognome:

ultimo cognome: verdi
```

