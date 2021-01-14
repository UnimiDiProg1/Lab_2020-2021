# Righe 

Scrivere un programma che:
* legga da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote ("")), terminando la lettura quando, premendo la combinazione di tasti Ctrl+D, viene inserito da **standard input** l'indicatore End-Of-File (EOF);
* utilizzando una funzione ricorsiva, stampi a video le righe lette (dalla prima letta all'ultima letta).

Oltre alla funzione `main()` di seguito riportata, deve essere definita ed utilizzata la sola funzione ricorsiva `LeggiEStampa(scanner *bufio.Scanner)` che riceve in input, nel parametro `scanner`, un'instanza del tipo `bufio.Scanner` inizializzata al fine di permettere la lettura da **standard input**. 

Per verificare il corretto funzionamento del programma, effettuare la redirezione dello standard output sul file `righe.txt`.  

```go
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	LeggiEStampa(scanner)
}
```


##### Esempio d'esecuzione:
```text
$ go run righe.go 
Questa è una riga
Questa è una riga
Seconda riga in input  
Seconda riga in input
Ogni riga viene stampata immediatamente
Ogni riga viene stampata immediatamente


Anche le righe vuote sono stampate
Anche le righe vuote sono stampate
```