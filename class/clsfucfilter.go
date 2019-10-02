package main

import (  
    "fmt"
)

type student struct {  
    firstName string
    lastName  string
    grade     string
    country   string
}

func filter(s []student, f func(student) bool) []student {  
    var r []student
    for _, v := range s {
        if f(v) == true {
            r = append(r, v)
        }
    }
    return r
}

func main() {  
    s1 := student{
        firstName: "vicky",
        lastName:  "D",
        grade:     "O+",
        country:   "India",
    }
    s2 := student{
        firstName: "suji",
        lastName:  "Div",
        grade:     "O+",
        country:   "Vellore",
    }
    s := []student{s1, s2}
    f := filter(s, func(s student) bool {
        if s.grade == "B" {
            return true
        }
        return false
    })
    fmt.Println(f)
}
