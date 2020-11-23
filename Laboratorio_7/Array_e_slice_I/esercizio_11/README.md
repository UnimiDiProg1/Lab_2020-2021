# Gioco dei dadi

Scrivere un programma che simuli il gioco dei dadi tra un numero variabile di giocatori.
Il programma deve leggere da **standard input** il numero di giocatori, i loro nomi, e il numero di turni da giocare.
Il nome di ogni giocatore è una stringa senza spazi.

Ad ogni turno di gioco, il programma deve simulare due lanci di dado per ogni giocatore (usando la funzione `rand.Intn()` del package `math/rand`). Ogni turno è vinto dal giocatore la cui somma dei lanci è maggiore. In caso di parità vince il giocatore che ha giocato prima. L'ordine di gioco è uguale all'ordine di inserimento dei nomi dei giocatori.

Prima di terminare, il programma deve stampare un riepilogo con il vincitore di ogni turno.

*Suggerimento:* per inizializzare il generatore di numeri randomici e avere un risultato diverso ad ogni esecuzione, inserite l'istruzione
```go
rand.Seed(int64(time.Now().Nanosecond()))
```
all'interno della funzione `main()`.


```text
$ go run dadi.go
Inserisci il numero di giocatori: 4
Inserisci i nomi dei 4 giocatori:
Mario
Luca
Sofia
Marta
Inserisci il numero di turni: 3
=== Turno 1 ===
        Giocatore Mario, lanci di valore: 3 e 5
        Giocatore Luca, lanci di valore: 2 e 5
        Giocatore Sofia, lanci di valore: 3 e 3
        Giocatore Marta, lanci di valore: 1 e 6
        Turno 1, vincitore: Mario
=== Turno 2 ===
        Giocatore Mario, lanci di valore: 5 e 2
        Giocatore Luca, lanci di valore: 2 e 6
        Giocatore Sofia, lanci di valore: 5 e 2
        Giocatore Marta, lanci di valore: 6 e 5
        Turno 2, vincitore: Marta
=== Turno 3 ===
        Giocatore Mario, lanci di valore: 3 e 2
        Giocatore Luca, lanci di valore: 1 e 2
        Giocatore Sofia, lanci di valore: 3 e 4
        Giocatore Marta, lanci di valore: 4 e 1
        Turno 3, vincitore: Sofia
Vittorie:
Vincitore turno 1: Mario
Vincitore turno 2: Marta
Vincitore turno 3: Sofia
```