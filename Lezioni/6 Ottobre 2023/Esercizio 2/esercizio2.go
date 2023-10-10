package main

import "fmt"

/* 
Data una data calcolare quanti giorni mancano a Natale.
*/

func main() {
	const giorniMese int = 30
	const giorniANatale int = (giorniMese * 12) + 25

	var mese, giorno int

	fmt.Print("Che data è? ")
	fmt.Scan(&mese, &giorno)

	giorniMancanti := giorniANatale - (mese * giorniMese) - giorno
	if (giorniMancanti > giorniANatale) {
		giorniMancanti += giorniMese * 12
	}

	fmt.Println("Mancano ", giorniMancanti, "giorni a Natale.")
}
