package main
import "fmt"

func main() {
	var x,y,z int

	x = 5
	y = (x+1) + z*3
	x = y*x
	z = x+y

	fmt.Println(x,z)
}