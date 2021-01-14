# Permutazioni

Scrivere un programma che legga da **standard input** una stringa senza spazi interamente definita da lettere dell'alfabeto inglese.  Ogni lettera è un elemento dell'insieme rappresentato dalla stringa (ogni lettera può comparire al più una volta nella stringa).

Una volta terminata la fase di lettura, il programma deve stampare a video tutte le possibili permutazioni semplici degli elementi dell'insieme rappresentato dalla stringa.

Oltre alla funzione `main()` di seguito riportata, deve essere definita ed utilizzata la sola funzione ricorsiva `func Permutazioni(insiemeResiduo string) []string` che riceve un valore `string` nel parametro `insiemeResiduo`, e restituisce un valore `[]string` in cui ogni elemento è una delle possibili permutazioni degli elementi dell'insieme rappresentato dalla stringa `insiemeResiduo`.

```go
func main() {

	insiemeResiduo := ""
    
    fmt.Scan(&insiemeResiduo)

	permutazioni := Permutazioni(insiemeResiduo)
	fmt.Printf("%d sottoinsiemi: %v\n", len(permutazioni), permutazioni)

}
```

##### Esempio d'esecuzione:
```text
$ go run permutazioni.go
ABC
6 permutazioni: [ABC ACB BAC BCA CAB CBA]

$ go run permutazioni.go
ciao
24 permutazioni: [ciao cioa caio caoi coia coai icao icoa iaco iaoc ioca ioac acio acoi aico aioc aoci aoic ocia ocai oica oiac oaci oaic]
```