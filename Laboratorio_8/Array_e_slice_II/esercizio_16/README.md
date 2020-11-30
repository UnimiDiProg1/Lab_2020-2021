# Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import (
	"fmt"
)

const Dimensione = 10

func main() {

	var a []int

	for i := 0; i < Dimensione; i++ {
		a = append([]int{i+1}, a...)
	}

	fmt.Println(a)
    
    b := make([]int, len(a))
    copy(b,a)
    fmt.Println(b)

    c := make([]int, Dimensione)
	copy(c[:], a[Dimensione/2:])
	fmt.Println(c)
}
```