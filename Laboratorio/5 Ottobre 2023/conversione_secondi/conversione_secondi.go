package main

import "fmt"

/*
Problema. Scrivere un programma Go conversione_secondi.go che converta un numero dato di secondi (fornito dall’utente) in giorni, ore, minuti, secondi senza mai usare l’operazione di sottrazione.

Esempio di esecuzione:
numero di secondi: 123456
g:h:m:s = 1:10:17:36

Nota: per stampare il risultato usare
fmt.Print(gg, ":", hh, ":", min, ":", sec, "\n")
(eventualmente sostituendo i nomi delle variabili con quelli che avete usato voi).
*/

func main() {
	const secondiMinuto int = 60
	const secondiOra int = secondiMinuto * 60
	const secondiGiorno int = secondiOra * 24

	var secondi int

	fmt.Print("numero di secondi: ")
	fmt.Scan(&secondi)

	fmt.Print("g:h:m:s = ", (secondi / secondiGiorno), ":", ((secondi % secondiGiorno) / secondiOra), ":", (((secondi % secondiGiorno) % secondiOra) / secondiMinuto), ":", (((secondi % secondiGiorno) % secondiOra) % secondiMinuto), "\n")
}