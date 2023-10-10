package main

import "fmt"

/*
Problema: Scrivere un programma Go centigradi_fahrenheit.go che converta in gradi Fahrenheit una temperatura espressa in gradi centigradi letta da standard input.
Si usino variabili di tipo float64.

Nota. Un grado Fahrenheit è 5./9. di un grado centigrado e 0 gradi centigradi corrisponde a 32
gradi Fahrenheit. Usare costanti (const) nella formula di conversione. Non usare direttamente i numeri (numeri magici). Attenzione alla differenza in
Go tra 5/9 e 5./9.
*/

func main() {
	const moltiplicatore float64 = 9./5.

	var temperaturaCentigradi float64
	var temperaturaFahrenheit float64

	fmt.Print("temperatura in centigradi? ")
	fmt.Scan(&temperaturaCentigradi)

	temperaturaFahrenheit = (temperaturaCentigradi * moltiplicatore) + 32

	fmt.Print(temperaturaCentigradi, "°C = ", temperaturaFahrenheit, " F")
}