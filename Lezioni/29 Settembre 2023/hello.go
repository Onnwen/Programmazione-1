package main
import "fmt"

func main() {
	var m, i int
	
	fmt.Print("Quante volte padre?   ") 
	fmt.Scan(&m)

	for i = 0; i < m; i++ {
		fmt.Println("Ciao")
	}
}