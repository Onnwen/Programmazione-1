package main
import "fmt"

func main() {
	var x,y,z int

	fmt.Scan(&x)
	z = x%10
	y = x%100
	fmt.Println((y-z)/10)
}