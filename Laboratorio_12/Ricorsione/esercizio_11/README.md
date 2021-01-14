# Divisione tra interi

Scrivere un programma che:
* legga da **riga di comando** due valori interi `dividendo` e `divisore`; 
* stampi a video il *quoziente* ed il *resto* della divisione tra interi `dividendo:divisore`, calcolati utilizzando una funzione ricorsiva.

*Suggerimento:* Il risultato della divisione tra interi `dividendo:divisore` può essere calcolato sottraendo ripetutamente il `divisore` al `dividendo` fino a che il primo non diventa maggiore del secondo. Il risultato della divisione tra interi `9:2` può essere per esempio calcolato come segue:
```text
9 - 2 = 7 (il 2 sta nel 9 una volta)
7 - 2 = 5 (il 2 sta nel 9 due volte)
5 - 2 = 3 (il 2 sta nel 9 tre volte)
3 - 2 = 1 (il 2 sta nel 9 quattro volte)
```  
Essendo 1 minore di 2, la procedura di calcolo si interrompe.
Il quoziente ed il resto della divisione `9:2` sono rispettivamente pari a `4` e `1`.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione ricorsiva `Dividi(dividendo, divisore int) (q, r int)` che riceve in input due valori `int`, rispettivamente nei parametri `dividendo` e `divisore`, e restituisce due valori `int` nelle variabili `q` ed `r` che rappresentano rispettivamente il quoziente ed il resto della divisione tra interi `dividendo:divisore`.

##### Esempio d'esecuzione:
```text
$ go run divisione.go 15 2
Quoziente = 7, Resto = 1

$ go run divisione.go 12 3
Quoziente = 4, Resto = 0

$ go run divisione.go 9 4 
Quoziente = 2, Resto = 1
```