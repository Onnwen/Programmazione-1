package main

/*
Scrivere un programma Go no_out.go che dichiara una variabile int e una costante = 10,
assegna alla variabile un valore doppio della costante, poi somma 1 alla variabile, ma non stampa niente.
*/

func main() {
	const b int = 10
	
	var a int = (b * 2) + 1
	_ = a
}