package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)
	ch <- "Vicky"
	ch <- "kumar"
	ch <- "smazz"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
