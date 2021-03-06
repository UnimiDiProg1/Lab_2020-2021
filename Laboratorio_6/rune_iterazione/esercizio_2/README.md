# Carte

Sapendo che al codice Unicode 127153 (associato alla rappresentazione in bit Unicode/UTF-8 `'\U0001F0B1'`) corrisponde il simbolo "asso di cuori", e che i codici successivi corrispondono alle carte successive (2 di cuori, 3 di cuori, ...), scrivere un programma che stampi tutte le carte da gioco dall'asso di cuori al 10 di cuori.

*Suggerimento:* Un carattere รจ una variabile di tipo `rune`, il cui valore รจ un intero corrispondente al codice Unicode del carattere. Le istruzioni equivalenti `var c rune = 127163` e `var c rune = '\U0001F0B1'` servono a definire la varibile `c` di tipo `rune` ed inizializzarla al valore "asse di cuori". Per stampare la carta da gioco "asse di cuori" si puรฒ utilizzare l'istruzione `fmt.Print(string(c))` o `fmt.Printf("%c",c)`.

```text
$ go run carte.go 
Simbolo: ๐ฑ - Codice numerico in base 10: 127153
Simbolo: ๐ฒ - Codice numerico in base 10: 127154
Simbolo: ๐ณ - Codice numerico in base 10: 127155
Simbolo: ๐ด - Codice numerico in base 10: 127156
Simbolo: ๐ต - Codice numerico in base 10: 127157
Simbolo: ๐ถ - Codice numerico in base 10: 127158
Simbolo: ๐ท - Codice numerico in base 10: 127159
Simbolo: ๐ธ - Codice numerico in base 10: 127160
Simbolo: ๐น - Codice numerico in base 10: 127161
Simbolo: ๐บ - Codice numerico in base 10: 127162
```