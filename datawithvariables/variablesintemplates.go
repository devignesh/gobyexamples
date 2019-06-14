package main
import (
	"log"
	"os"
	"text/template"

)

var v1 *template.Template
func init() {
	v1 := template.Must(template.ParseFiles("vicky.gohtml"))

}

func main() {
	err := v1.ExecuteTemplate(os.Stdout, "vicky.gohtml", "Vicky kumar")
	if err != nil {
		log.Fatalln(err)
	}
}