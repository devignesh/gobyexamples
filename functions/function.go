package main
import "fmt"
func plus(a int, b int) int {
	return a + b
}

func plusm(a, b, c int) int {
	return a + b + c
}

func main() {
	add := plus(2, 5)
	addm :=plusm(3, 5, 5)
	fmt.Println("value of add is:", add)
	fmt.Println("Value of addm is:", addm)
}