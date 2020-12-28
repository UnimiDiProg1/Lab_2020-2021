# Frazioni

Al fine di modellare l'entità matematica `frazione`, si definisca un nuovo tipo `Frazione` come una struttura avente due campi `numeratore` e `denominatore` di tipo `int`.

Implementare le funzioni:

* `NuovaFrazione(numeratore, denominatore int) *Frazione` che restituisce una nuova istanza del tipo `Frazione` inizializzata in base ai valori dei parametri `numeratore` e `denominatore`;
* `String(f Frazione) string` che riceve in input un'instanza del tipo `Frazione` nel parametro `f` e restituisce un valore `string` che corrisponde alla rappresentazione `string` di `f` nel formato `numeratore/denominatore`;

Verificare il corretto funzionamento delle funzioni utilizzando le unit di test allegate.

*Suggerimento:* per creare una stringa formattata potete utilizzare la funzione `fmt.Sprintf()` del package `fmt`. La funzione si comporta come la funzione `fmt.Printf()` ma, invece di stampare a video l'output, restituisce una `string` contenente l'output.