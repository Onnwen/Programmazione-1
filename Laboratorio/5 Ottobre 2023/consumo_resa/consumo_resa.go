package main

import "fmt"

/*
Problema. Scrivere un programma Go consumo_resa.go che calcola il consumo medio e la resa di un motore date la distanza totale percorsa (in km) e la quantità di carburante utilizzata (in litri). I valori in input sono di tipo float64.

Nota. Il consumo medio di carburante si esprime in l/km ed è la quantità di carburante che occorre in media per percorrere un km di strada. La resa di un motore è data dalla distanza percorsa in media con un litro di carburante e si esprime in km/l.

Esempio di esecuzione:
distanza percorsa (in km): 50
quantità di carburante utilizzata (in l): 10
consumo medio: 0.2 l/km
resa media: 5 km/l
*/

func main() {
	var distanzaPercorsa float64
	var litriCarburanteUtilizzati float64

	fmt.Print("Inserire la distanza percorsa (in km) e la quantità di carburante utilizzata (in litri): ")
	fmt.Scan(&distanzaPercorsa, &litriCarburanteUtilizzati)

	fmt.Println("distanza percorsa (in km): ", distanzaPercorsa)
	fmt.Println("quantità di carburante utilizzata (in l): ", litriCarburanteUtilizzati)
	fmt.Println("consumo medio: ", litriCarburanteUtilizzati / distanzaPercorsa, " l/km")
	fmt.Println("resa media: ", distanzaPercorsa / litriCarburanteUtilizzati, " km/l")
}