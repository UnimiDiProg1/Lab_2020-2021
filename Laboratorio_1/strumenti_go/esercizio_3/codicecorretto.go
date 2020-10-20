package main

import "fmt" // il package fmt va tra virgolette

func main() { // il main necessita delle parentesi tonde

	var x = 10
	var y = 15

	sum := x + y // la variabile sum non è stata dichiarata

	fmt.Println("La somma è ", sum) // la parentesi tonda non era chiusa

}

// } una parentesi graffa di troppo
