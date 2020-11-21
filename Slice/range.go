package main

import (
	"fmt"
)

func main() {
	x := []int{23, 24, 45, 56, 67}

	for ind, val := range x {
		fmt.Println(ind, val)
	}
}
