package main

import "fmt"

/*
Dato un orario (ore e minuti) calcolare quanti minuti mancano a mezzanotte.
*/

func main() {
	const oreGiornata int = 23
	const minutiOra int = 60

	var ore, minuti int

	fmt.Print("Che ora è? ")
	fmt.Scan(&ore, &minuti)


	fmt.Print("Mancano ", (oreGiornata - ore) * 60 + (minutiOra - minuti), " minuti")
}