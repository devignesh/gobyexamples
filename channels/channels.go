package main
import "fmt"
func main() {
		message := make(chan string)
		go func() { message <-"vicky"} ()
		msg := <-message
		fmt.Println(msg)
}