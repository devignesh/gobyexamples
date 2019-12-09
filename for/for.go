package main 

import "fmt"
func main() {
	i:=1
	for i<=10000000000000000 {
		i=i*2
		fmt.Println(i)

		if i%2 == 0{
			fmt.Println("EVEN NO")
		}else {
			fmt.Println("ODD NO")
		}

		i++;
	}
}