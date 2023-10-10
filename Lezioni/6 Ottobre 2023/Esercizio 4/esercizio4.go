package main

import "fmt"

/* 
Dato totale e aliquota, calcola l'imponibile.
*/

func main() {
	var totale, aliquota float64

	fmt.Print("Inserire totale: ")
	fmt.Scan(&totale)
	fmt.Print("Inserire aliquota: ")
	fmt.Scan(&aliquota)

	imponibile := (100 * totale) / (100 + aliquota)
	iva := totale - imponibile
	fmt.Println("\nImponibile: ", imponibile, "€")
	fmt.Println("IVA: ", iva, "€")
	fmt.Println("Totale: ", totale, "€")
}