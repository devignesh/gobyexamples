package main
import "fmt"
func value() (int, int,) {
	return 5, 7
}

func main() {
	a, b := value()
	fmt.Println(a)
	fmt.Println(b)
	_, c :=value()
	fmt.Println(c)
}