# Riduzione di una frazione ai minimi termini

Si consideri il tipo `Frazione` dell'esercizio **5 Frazioni**.

Si modifichi il programma in modo tale che:
* legga da **riga di comando** due numeri interi `n` e `d`;
* utilizzi `n` e `d` per inizializzare una `Frazione` utilizzando `n` come numeratore e `d` come denominatore;
* riduca ai minimi termini la frazione; 
* stampi a video, nel formato `numeratore/denominatore`, la frazione ridotta ai minimi termini.

Oltre alla funzione `main()` e alle funzioni implementate nell'esercizio **5 Frazioni**, devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `Riduci(f *Frazione)` che riceve in input un'instanza del tipo `Frazione` nel parametro `f` e, se necessario, modifica opportunamente il valore dei campi `f.numeratore` e `f.denominatore` affinché `f` rappresenti una frazione ridotta ai minimi termini.

Verificare il corretto funzionamento delle funzioni utilizzando le unit di test allegate.

##### Esempio d'esecuzione:
```text
$ go run riduci.go 10 10
1/1

$ go run riduci.go 34 18
17/9

$ go run riduci.go 12 36
1/3
```