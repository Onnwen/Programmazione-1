package main

import "fmt"

/*
Date le altezze in cm di du epersone, calcolare la media.
*/

func main() {
	var altezza1, altezza2 int

	fmt.Print("Altezza 1: ")
	fmt.Scan(&altezza1)
	fmt.Print("Altezza 2: ")
	fmt.Scan(&altezza2)

	fmt.Println("Altezza media: ", float64(altezza1 + altezza2) / 2.)
}