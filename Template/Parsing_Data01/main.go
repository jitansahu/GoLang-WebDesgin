package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	if err := tpl.ExecuteTemplate(os.Stdout, "index.html", 42); err != nil {
		fmt.Println(err)
	}
}
