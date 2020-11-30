# Date (2)

Scrivere un programma che:
* legga da **standard input** un testo su più righe;
* termini la lettura quando viene inserita da standard input una riga vuota ( "" ).

Ogni riga di testo è una una stringa senza spazi che codifica una data in uno dei seguenti possibili formati:

1. aaaa/m/g
2. aaaa/m/gg
3. aaaa/mm/g
4. aaaa/mm/gg
5. g/m/aaaa
6. gg/m/aaaa
7. g/mm/aaaa
8. gg/mm/aaaa

Una volta terminata la fase di lettura, il programma deve stampare la codifica nel formato aaaa-mm-gg delle date lette. In particolare, le date devono essere stampate in ordine cronologico (dalla più antica alla più recente).

Oltre alla funzione `main()` devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `daInvertire(s string) bool` che riceve in input un valore `string` nel parametro `s` in cui è codificata una data nel formato **aaaa/m/g**, **aaaa/m/gg**, **aaaa/mm/g**, **aaaa/mm/gg**, **g/m/aaaa**, **gg/m/aaaa**, **g/mm/aaaa**, o **gg/mm/aaaa** e restituisce `true` se in `s` è codificata una data nel formato **g/m/aaaa**, **gg/m/aaaa**, **g/mm/aaaa**, o **gg/mm/aaaa**, `false` altrimenti.
* una funzione `Inverti(s string) string` che riceve in input un valore  `string` nel parametro `s` e restituisce un valore `string` in cui il primo carattere è l'ultimo che definisce `s`, il secondo carattere è il penultimo che definisce `s`, ...  e l'ultimo carattere è il primo che definisce `s`.

*Suggerimento:* Una volta codificate nel formato **aaaa-mm-gg**, le date possono essere ordinate cronologicamente utilizzando la funzione `sort.Strings(a []string)` del package `sort`.

##### Esempio d'esecuzione:

```text
$ cat test
2019/04/03
2019/6/4
2019/07/5
2019/4/05
5/5/2019
05/4/2019
7/05/2019
07/09/2019
07/09/2018

$ go run date.go < test      
2018-09-07
2019-04-03
2019-04-05
2019-04-05
2019-05-05
2019-05-07
2019-06-04
2019-07-05
2019-09-07
```