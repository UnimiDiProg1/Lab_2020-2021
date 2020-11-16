# Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import "fmt"

func main() {

	s := "ciao, come va?"
    /* s è interamente definita da caratteri considerati nello standard US-ASCII */

    fmt.Println(s[6:10] + string(s[len(s)-1]))
	fmt.Println(s[:5] + s[len(s)-4:])
	
}
```