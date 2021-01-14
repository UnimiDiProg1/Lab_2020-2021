# Sottoinsiemi

Scrivere un programma che legga da **standard input** una stringa senza spazi interamente definita da lettere dell'alfabeto inglese.  Ogni lettera è un elemento dell'insieme rappresentato dalla stringa (ogni lettera può comparire al più una volta nella stringa).

Una volta terminata la fase di lettura, il programma deve stampare a video tutti i possibili sottoinsiemi dell'insieme rappresentato dalla stringa (insieme vuoto escluso).

Oltre alla funzione `main()` di seguito riportata, deve essere definita ed utilizzata la sola funzione ricorsiva `func Sottoinsiemi(insiemeResiduo string) []string` che riceve un valore `string` nel parametro `insiemeResiduo`, e restituisce un valore `[]string` in cui ogni elemento è uno dei possibili sottoinsiemi dell'insieme rappresentato dalla stringa `insiemeResiduo`.

```go
func main() {

	insiemeResiduo := ""
    
    fmt.Scan(&insiemeResiduo)

	sottoinsiemi := Sottoinsiemi(insiemeResiduo)
	fmt.Printf("%d sottoinsiemi: %v\n", len(sottoinsiemi), sottoinsiemi)

}
```

##### Esempio d'esecuzione:
```text
$ go run sottoinsiemi.go 
ABC
7 sottoinsiemi: [C B BC A AC AB ABC]

$ go run sottoinsiemi.go
ciao
15 sottoinsiemi: [o a ao i io ia iao c co ca cao ci cio cia ciao]
```