# Trova l'errore

Questo programma contiene degli errori. Corregere gli errori ed eseguire il programma.

```go
package main

import "fmt"

func test(x int) (y int, z int) {
	y = x + 1
	z = x + 2
	return
}

func main() {
	var a, b int
	a, b = test(10)
	fmt.Println(a, b)
}
```