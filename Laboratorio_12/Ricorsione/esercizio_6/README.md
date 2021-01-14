# Righe invertite

Scrivere un programma che:
1. legga da **standard input** una stringa senza spazi;
2. inverta la stringa letta utilizzando una funzione ricorsiva e stampi a video la stringa ottenuta.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione ricorsiva `InvertiStringa(s string) string` che riceve in input un valore `string` nel parametro `s` e restituisce un valore `string` in cui il primo carattere è
l'ultimo che definisce `s`, il secondo carattere è il penultimo che definisce `s`, ... e l'ultimo carattere è il
primo che definisce `s `.

##### Esempio d'esecuzione:
```text
$ go run inverti_stringa.go 
programmazione
enoizammargorp

$ go run inverti_stringa.go
laboratorio
oirotarobal
```