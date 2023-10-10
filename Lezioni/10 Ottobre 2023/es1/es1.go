package main

import "fmt"

/*
Leggere due frazioni (numeratore e denominatore) e stabilire quale sia la più piccola
*/

func main() {
	var numeratore1, denominatore1, numeratore2, denominatore2 int

	fmt.Print("Inserire la prima frazione: ")
	fmt.Scan(&numeratore1, &denominatore1)
	fmt.Print("Inserire la seconda frazione: ")
	fmt.Scan(&numeratore2, &denominatore2)


	fmt.Print(numeratore1, "/", denominatore1)

	p1 := numeratore1 * denominatore2
	p2 := numeratore2 * denominatore1

	if p1 > p2 {
		fmt.Print(" > ")
	} else if p1 < p2 {
		fmt.Print(" < ")
	} else {
		fmt.Print(" = ")
	}

	fmt.Println(numeratore2, "/", denominatore2)
}