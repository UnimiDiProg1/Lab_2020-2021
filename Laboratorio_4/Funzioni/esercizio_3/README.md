# Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import "fmt"

func test(x int) (y int) {
	y = x * 10
	x, y = y, x
	return
}

func main() {
	var a, b int = 10, 5
	b = test(a)
	fmt.Println(a, b)
}
```