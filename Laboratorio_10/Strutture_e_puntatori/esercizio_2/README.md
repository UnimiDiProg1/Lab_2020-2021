# Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	var a, b = 15, 20
	f(&a, &b)
	fmt.Println(a, b)
}
func f(x, y *int) {
	*x, *y = *y, *x
}
```
