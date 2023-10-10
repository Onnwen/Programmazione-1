package main

import "fmt"

/*
Problema. Scrivere un programma Go intersezione_rette.go che, date le equazioni di due rette, stabilisce in che punto si intersecano. I dati in input sono di tipo float64.

Nota 1. L’equazione della retta è: y = mx + q 
dove m è il coefficiente angolare della retta e q l'intercetta dell'asse y.

Nota 2. Dal punto di vista matematico, per determinare il punto di intersezione tra le due rette, occorre risolvere un sistema di primo grado.

Esempio di esecuzione:
retta 1: m e q? 1 4
retta 2: m e q? 2 6
intersezione in (-2,2)
*/

func main() {
	var m1, q1 float64
	var m2, q2 float64

	fmt.Print("retta 1: m e q? ")
	fmt.Scan(&m1, &q1)

	fmt.Print("retta 2: m e q? ")
	fmt.Scan(&q2, &q2)

	var x float64 = ((q2 - q1) / (m1 - m2))
	var y float64 = (m1 * x + q1)
	fmt.Print("intersezione in (", x, ",", y, ")")
}