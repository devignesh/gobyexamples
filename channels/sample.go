package main

import (
	"fmt"
	"runtime"
)

func sample() {
	fmt.Println(runtime.GOARCH)
}
func main() {
	sample()
}
