package main
import "fmt"
func main() {
	message := make(chan string, 2)
	message <- "Vicky"
	message <- "MCA"
	fmt.Println(<-message)
	fmt.Println(<-message)
}