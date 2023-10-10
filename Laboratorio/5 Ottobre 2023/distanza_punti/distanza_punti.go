package main

import (
	"fmt"
	"math"
)

/*
Problema: Scrivere un programma Go distanza_punti.go che calcola la distanza tra due punti nel piano cartesiano.

Annotazioni: Siano P1 = (x1, y1) e P2 = (x2, y2). La formula della distanza tra due punti P1 e P2 è:
radiceQuadrata( (x2 − x1)^2 + (y2 − y1)^2 )
dove ^2 indica l'elevamento alla seconda.
Siano x1, y1, x2, y2 tutti valori float64.

Nota. La funzione per calcolare la radice quadrata si chiama Sqrt ed è messa a disposizione nel package "math". Per calcolare invece il quadrato z^2 di un numero z si usi z*z.

Esempio di esecuzione:
x e y del primo punto: 5 8
x e y del secondo punto: 7 11
La distanza tra i due punti è: 3.605551275463989
*/

func main() {
	var x1, y1 float64
	var x2, y2 float64

	fmt.Print("x e y del primo punto: ")
	fmt.Scan(&x1, &y1)

	fmt.Print("x e y del secondo punto: ")
	fmt.Scan(&x2, &y2)

	fmt.Println("La distanza tra i due punti è: ", math.Sqrt(((x2 - x1) * (x2 - x1)) + (y2 - y1) * (y2 - y1)))
}