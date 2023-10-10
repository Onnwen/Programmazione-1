package main

import "fmt"

/*
Convertire la media da 30esimi in 110esimi
*/

func main() {
	var input int

	fmt.Print("Inserire media in 30esimi: ")
	fmt.Scan(&input)

	fmt.Print("Media in 110esimi: ", (input * 110) / 30)
}