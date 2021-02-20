package main

import (
	"fmt"
)

type vehicle interface {
	accelerate()
}

func foo(v vehicle) {
	fmt.Println("foo ", v)

}

type car struct {
	model string
	color string
}

func (c car) accelerate() {
	fmt.Println("Accelrating?")

}

type toyota struct {
	model string
	color string
	speed int
}

func (t toyota) accelerate() {
	fmt.Println("I am toyota, I accelerate fast?")
}
func main() {
	c1 := car{"suzuki", "blue"}
	t1 := toyota{"Toyota", "Red", 100}
	c1.accelerate()
	t1.accelerate()
	foo(c1)
	foo(t1)
}
