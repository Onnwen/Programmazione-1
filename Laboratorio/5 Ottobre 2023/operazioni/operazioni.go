package main

/*
Problema: Completare il programma Go operazioni.go in modo che, dati due numeri float64 da standard
input, esegua addizione, sottrazione, moltiplicazione e divisione dei due numeri e stampi i risultati ottenuti.

Nota. Con la conversione in binario c’è la possibilità, da tenere sempre presente, che si
introducano errori di approssimazione per i float64 e di overflow per gli int (lo vedrete più avanti)
*/

import "fmt"

func main() {
	//dichiarare qui due variabili x1 e x2 di tipo float64
	var x1 float64
	var x2 float64

	fmt.Print("fornire due numeri (float): ")
	fmt.Scan(&x1, &x2)
	
	//inserire qui le istruzioni per svolgere le operazioni
	//salvando i risultati con degli short assignment
	// (i nomi delle variabili da utilizzare per i risultati sono quelli usati nelle istruzioni di stampa qui sotto, cioè somma, differenza, prodotto, quoziente, )
	somma := x1 + x2
	differenza := x1 - x2
	prodotto := x1 * x2
	quoziente := x1 / x2

	fmt.Println(x1, "+", x2, "=", somma)
	fmt.Println("differenza =", differenza)
	fmt.Println("il prodotto di", x1, "e", x2, "dà", prodotto)
	fmt.Println(x1, "/", x2, "=", quoziente)
}
