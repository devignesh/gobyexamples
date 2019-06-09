package main
import "fmt"
func main() {
	message := make(chan string, 2)
	message <- "Vicky"
	message <- "MCA"
	fmt.Pri(<-message)
	fmt.Println(<-message)
}