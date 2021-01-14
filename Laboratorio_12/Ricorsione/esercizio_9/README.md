# Testo invertito

Scrivere un programma che:
* legga da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote ("")), terminando la lettura quando, premendo la combinazione di tasti Ctrl+D, viene inserito da **standard input** l'indicatore End-Of-File (EOF);
* ristampi il testo letto dall’ultimo carattere al primo.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione ricorsiva `InvertiStringa(s string) string` che riceve in input un valore `string` nel parametro `s` e restituisce un valore `string` in cui il primo carattere è
l'ultimo che definisce `s`, il secondo carattere è il penultimo che definisce `s`, ... e l'ultimo carattere è il
primo che definisce `s `;
* una funzione ricorsiva `LeggiEStampa(scanner *bufio.Scanner)` che riceve in input, nel parametro `scanner`, un'instanza del tipo `bufio.Scanner` inizializzata al fine di permettere la lettura da **standard input**.

Per verificare il corretto funzionamento del programma, effettuare la redirezione dello standard output sul file `testo_invertito.txt`.  

##### Esempio d'esecuzione:
```text
$ go run inverti_testo.go 
Questa è una riga
Seconda riga in input
 
Anche le righe vuote contano
onatnoc etouv ehgir el ehcnA

tupni ni agir adnoceS
agir anu è atseuQ
```