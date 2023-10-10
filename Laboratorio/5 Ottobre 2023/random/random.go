package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Problema: Scrivere un programma Go random.go che generi un numero random int tra 0 e 100 e lo stampi.
Nota. Vedere la documentazione della funzione Intn del package "math/rand". Per
generare un numero diverso a ogni esecuzione importare il package "time" e invocare la
funzione rand.Seed(time.Now().Unix()) prima di invocare rand.Intn.
*/

func main() {
	rand.Seed(time.Now().Unix())

	fmt.Println("Numero casuale: ", rand.Intn(100))
}