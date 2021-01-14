# Frazione continua

**Definizione**: Sia a<sub>0</sub> un numero intero qualsiasi e a<sub>1</sub>, a<sub>2</sub>, ... , a<sub>n</sub> numeri interi positivi. 
La notazione [a<sub>0</sub>, a<sub>1</sub>, ... , a<sub>n</sub>] indica la frazione continua

<img src="https://latex.codecogs.com/png.latex?\fn_phv&space;a_0&space;&plus;&space;\frac{1}{a_1&space;&plus;&space;\frac{1}{a_2&space;&plus;&space;\frac{1}{a_3&space;&plus;&space;\dots}}}" title="[a_0; a_1, a_2, \dots, a_n] = a_0 + \frac{1}{a_1 + \frac{1}{a_2 + \frac{1}{a_3 + \dots}}}" />

Ad esempio: 

<img src="https://latex.codecogs.com/png.latex?\fn_phv&space;[-1,&space;5,&space;2,&space;4]&space;=&space;-1&space;&plus;&space;\frac{1}{5&space;&plus;&space;\frac{1}{2&space;&plus;&space;\frac{1}{4}&space;}&space;}" title="[-1, 5, 2, 4] = -1 + \frac{1}{5 + \frac{1}{2 + \frac{1}{4} } }" />

e allo stesso modo:

<img src="https://latex.codecogs.com/png.latex?\fn_phv&space;[-7,&space;3,&space;5,&space;7,&space;9]&space;=&space;-7&space;&plus;&space;\frac{1}{3&space;&plus;&space;\frac{1}{5&space;&plus;&space;\frac{1}{7&space;&plus;&space;\frac{1}{9}}&space;}&space;}" title="[-7; 3, 5, 7, 9] = -7 + \frac{1}{3 + \frac{1}{5 + \frac{1}{7 + \frac{1}{9}} } }" />

Scrivere un programma che:
* legga da **riga di comando** una sequenza di numeri interi a<sub>0</sub> a<sub>1</sub> a<sub>2</sub> ...  a<sub>n</sub> (con a<sub>i</sub> > 0 per 1 <= i <= n);
* stampi a video il valore reale corrispondente alla frazione continua [a<sub>0</sub>, a<sub>1</sub>, ... , a<sub>n</sub>], calcolato utilizzando una funzione ricorsiva.


Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione ricorsiva `ValoreFrazione(sl []int) float64` che riceve in input un valore `[]int` nel parametro `sl` e restituisce il valore `float64` corrispondente alla frazione continua `[sl[0], sl[1], ... , sl[n-1]]`, dove `n` è pari a `len(sl)`.

##### Esempio d'esecuzione:
```text
$ go run frazioneContinua.go -1 2 2
Frazione continua: -0.6

$ go run frazioneContinua.go -1 2  
Frazione continua: -0.5

$ go run frazioneContinua.go -1 4 6  
Frazione continua: -0.76

$ go run frazioneContinua.go 2 4 6  
Frazione continua: 2.24

$ go run frazioneContinua.go -1 4 6 5
Frazione continua: -0.7596899224806202
```
