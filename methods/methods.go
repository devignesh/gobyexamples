package main

import (  
    "fmt"
)

type Employee struct {  
    name     string
    salary   int
    currency string
}


func (e Employee) displaySalary() {  
    fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func main() {  
    emp1 := Employee {
        name:     "vicky",
        salary:   13000,
        currency: "$",
    }
    emp1.displaySalary() 
}
