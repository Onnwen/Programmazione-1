package main

import "fmt"

/* 
Dato un imponibile e un'aliquota percentuale, calcola l'IVA.
*/

func main() {
	var imponibile, aliquota float64

	fmt.Print("Inserire imponibile: ")
	fmt.Scan(&imponibile)
	fmt.Print("Inserire aliquota: ")
	fmt.Scan(&aliquota)

	iva := imponibile * (aliquota / 100)
	totale := imponibile + iva
	fmt.Println("\nImponibile: ", imponibile, "€")
	fmt.Println("IVA: ", iva, "€")
	fmt.Println("Totale: ", totale, "€")
}