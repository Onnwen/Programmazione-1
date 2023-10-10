package main

import "fmt"

/*
Problema: Scrivere un programma Go angolo_triangolo.go che, date in input le ampiezze in
gradi (float64) di due angoli di un triangolo, determini l’ampiezza del terzo angolo.
Nota. La somma degli angoli di un triangolo è sempre 180 gradi, è quindi il caso di utilizzare una costante per questo valore.
*/

func main() {
	const sommaAngoli float64 = 180

	var ampiezza1 float64
	var ampiezza2 float64

	fmt.Print("Ampiezza angolo n°1: ")
	fmt.Scan(&ampiezza1)
	fmt.Print("Ampiezza angolo n°2: ")
	fmt.Scan(&ampiezza2)

	fmt.Print("Ampiezza angolo n°3: ", sommaAngoli - (ampiezza1 - ampiezza2))
}