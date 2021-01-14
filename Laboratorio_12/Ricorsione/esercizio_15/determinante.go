package main

import (
	"fmt"
	"math"
)

func determinante(matrice [][]float64) (risultato float64) {

	if len(matrice) == 2 { //caso base

		risultato = matrice[0][0]*matrice[1][1] - matrice[0][1]*matrice[1][0]

	} else { //riduzione a caso più semplice

		for j := 0; j < len(matrice); j++ { //per ogni colonna

			//creazione di una sottomatrice più piccola
			sottomatrice := make([][]float64, len(matrice)-1) //inizializza la sottomatrice

			for i := 1; i < len(matrice); i++ { //per ogni riga (tranne la prima)
				//crea la riga della sottomatrice
				sottomatrice[i-1] = append(sottomatrice[i-1], matrice[i][:j]...)   //append di elementi prima di j
				sottomatrice[i-1] = append(sottomatrice[i-1], matrice[i][j+1:]...) //append di elementi dopo j
			}

			//ricorsione
			risultato += matrice[0][j] * math.Pow(-1, float64(j)) * determinante(sottomatrice)
		}
	}

	//fmt.Println("determinante per sottomatrice",matrice,"=",risultato)
	return risultato
}

func main() {
	matrice := [][]float64{
		{1, 2, 3},
		{-2, -1, -3},
		{0, -4, 1},
	}

	fmt.Println(determinante(matrice))
}
