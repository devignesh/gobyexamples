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
	fmt.Println(runtime.GOOS)
	v := "vicky"
	fs := []byte(v)
	fmt.Println(fs)
	fmt.Printf("%T\n", fs)

	for i := 0; i < len(v); i++ {
		fmt.Println("%#U", v[i])
	}
}
