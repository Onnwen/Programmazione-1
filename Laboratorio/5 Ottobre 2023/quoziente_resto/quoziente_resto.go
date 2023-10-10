package main

/*
Problema: Scrivere un programma Go quoziente_resto.go che legge da standard input un
dividendo e un divisore (interi), calcola il quoziente e il resto e li stampa.
Nota. L’operatore per la divisione (/) tra interi calcola la parte intera del risultato; l’operatore
per il resto della divisione è %.
*/

import "fmt"

func main() {
	//dichiarare qui due variabili dividendo e divisore di tipo int
	var dividendo int 
	var divisore int

	fmt.Println("fornire due numeri (int) di cui si vuole calcolare quoziente e resto")
	fmt.Print("dividendo: ")
	fmt.Scan(&dividendo)
	fmt.Print("divisore: ")
	fmt.Scan(&divisore)
	
	//inserire qui le istruzioni per svolgere le operazioni
	//salvando i risultati con degli short assignment
	// (per i nomi delle variabili da utilizzare per i risultati, vedere le istruzioni di stampa qui sotto)
	quoziente := dividendo / divisore
	resto := dividendo % divisore

	fmt.Println("quoziente =", quoziente)
	fmt.Println("resto =", resto)
}
