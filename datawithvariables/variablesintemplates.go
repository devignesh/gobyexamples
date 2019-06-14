package main
import (
	"log"
	"os"
	"text/template"

)

var tpl *template.Template

func init() {
	tpl := template.Must(template.ParseFiles("vicky.gohtml"))

}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "vicky.gohtml", "Vicky kumar")
	if err != nil {
		log.Fatalln(err)
	}
}