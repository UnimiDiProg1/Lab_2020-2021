# Determinante

Scrivere un programma che calcoli il determinante di una matrice quadrata secondo il primo [Teorema di Laplace](https://it.wikipedia.org/wiki/Teorema_di_Laplace).

Il caso base è il calcolo per il determinante di una matrice 2x2:
![Determinante di una matrice 2x2](https://wikimedia.org/api/rest_v1/media/math/render/svg/0b9760e168202a5f91a279523fcd1e738e2755ff)

Invece, per matrici quadrate di dimensione superiore il determinante si calcola mediante la seguente procedura:

- si seleziona una riga i della matrice (per semplicità scegliamo la riga 0).
- per ciascuna colonna j si somma al risultato il prodotto dell'elemento ![m_{i,j}](https://wikimedia.org/api/rest_v1/media/math/render/svg/cd6f1bb2d6548dca472922bcbcb77e7ad3a5b4df) per ![-1^(i+j) det(M_{i,j)](https://wikimedia.org/api/rest_v1/media/math/render/svg/d46a48317bfaf753cd878b5b46fb625f1ff04b53) dove ![M_{i,j}](https://wikimedia.org/api/rest_v1/media/math/render/svg/6136d919b3f10cc58cc023e05a5c36b2f44e111f) è la sottomatrice (di dimensione n − 1) che si ottiene dalla matrice cancellando la i-esima riga e la j-esima colonna.

Ad esempio, per la matrice ![Matrice 3x3](https://wikimedia.org/api/rest_v1/media/math/render/svg/58e23b99b9b8aaac4e120a7861dec2924e6a9894) considerendo la prima riga per ciascuna colonna abbiamo:

![Elemento 1](https://wikimedia.org/api/rest_v1/media/math/render/svg/a11664e832f0e7605c2bc10cec64058662fa6bb2)
![Elemento 2](https://wikimedia.org/api/rest_v1/media/math/render/svg/850c43861162407e476f867a77125c52b7c54087)
![Elemento 3](https://wikimedia.org/api/rest_v1/media/math/render/svg/0ebeabd4510faad713b8c807a9323f2f624a108d)

Pertanto il determinante è ![Determinante](https://wikimedia.org/api/rest_v1/media/math/render/svg/cced1aa208435c489ab182dec2d6d85ac6a9751d)

Per semplicità, utilizzare il seguente codice main:

```go
func main() {
	matrice := [][]float64{
		{ 1, 2, 3},
		{-2,-1,-3},
		{ 0,-4, 1},
	}
	
	fmt.Println(determinante(matrice))
}
```

##### Esempio d'esecuzione:
```text
$ go run determinante.go
15
```
