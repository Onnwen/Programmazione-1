package main

import "fmt"

/*
Problema: Scrivere un programma Go conversione_giorni.go che converta un numero specificato di giorni (letto da tastiera) in anni, settimane, giorni senza mai usare l’operazione di sottrazione. Si considerino tutti gli anni di 365 giorni.

Nota. Dichiarare costanti (const) per il numero di giorni (365) in un anno e il numero di giorni in una
settimana (7). Non usare direttamente i numeri (numeri magici).

Esempio di esecuzione:
numero di giorni da convertire?
1329
1329 giorni = 3 anni + 33 settimane + 3 giorni
*/

func main() {
	const giorniAnno int = 365
	const giorniSettimana int = 7

	var giorni int

	fmt.Print("numero di giorni da convertire? ")
	fmt.Scan(&giorni)

	fmt.Print(giorni, " giorni = ", (giorni / giorniAnno), " anni + ", ((giorni % giorniAnno) / giorniSettimana), " settimane + ", ((giorni % giorniAnno) % giorniSettimana), " giorni")
}