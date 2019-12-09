package main
import "fmt"
func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println("vicky")
	} 
	
	j := 1
	for j <= 3 {
		fmt.Println(j)
		j = j + 1
	}

	for {
        fmt.Println("Vignesh")
        break
    }
	for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}
