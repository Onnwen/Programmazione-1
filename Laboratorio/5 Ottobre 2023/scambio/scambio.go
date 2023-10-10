package main

import "fmt"

/*
Problema: Completare il programma Go scambio.go in modo che, dati in input due valori di tipo int,

stampi l'eventuale errore nella lettura,
stampi i due valori nell’ordine in cui sono stati forniti,
scambi i due valori (senza appoggiarsi ad altre variabili),
esegua la stessa istruzione di stampa (per mostrare che i valori sono stati effettivamente scambiati tra le due variabili).
*/

func main() {
	//dichiarare qui due variabili n1 e n2 di tipo int
	var n1 int
	var n2 int

	fmt.Print("fornire due numeri (int): ")
	_, err := fmt.Scan(&n1, &n2)
	fmt.Println(err) //stampo l'errore , se è nil, vuol dire che non si sono verificati errori 
	//inserire qui l'istruzione per fare lo scambio

	fmt.Println("Numero n°1: ", n1, "\nNumero n°2: ", n2)

	n1, n2 = n2, n1
	
	fmt.Println("---\nNumero n°1: ", n1, "\nNumero n°2: ", n2)
}
