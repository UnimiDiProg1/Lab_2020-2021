# Statistiche testo

Scrivere un programma che: 
* legga da **standard input** un testo su più righe;
* termini la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito da **standard input** l'indicatore End-Of-File (EOF);
* stampi a video le seguenti statistiche relative al testo letto:
  1. il numero di parole presenti nel testo;
  2. la lunghezza media delle parole presenti nel testo.

**Nota:** una parola è una sequenza di caratteri consecutivi rappresentanti delle lettere. Numeri, punteggiatura e caratteri di spaziatura intervallano parole diverse.

Oltre alla funzione `main()`, il programma deve definire e utilizzare le seguenti funzioni:
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe terminato dall'indicatore EOF (`CTRL+D`), restituendo un valore `string` in cui è memorizzato il testo letto;
* una funzione `StatisticheParole(s string) (int, int)` che riceve in input un valore `string` nel parametro `s` e restituisce due valori `int` pari rispettivamente al numero di parole presenti in `s` e alla loro lunghezza totale.

*Suggerimento:* per sapere se un carattere rappresenta una lettera utilizza la funzione `unicode.IsLetter()` del package `unicode`. Utilizza `go doc unicode.IsLetter` per scoprire il suo funzionamento.

##### Esempio d'esecuzione:

```text
$ go run statistiche.go    
Inserisci un testo su più righe (termina con Ctrl+D):
Testo di prova                      

su cui calcolare le statistiche.
Test 01: prova.prova0prova
Statistiche:
Numero parole: 12
Lunghezza media: 4.833333333333333
```
