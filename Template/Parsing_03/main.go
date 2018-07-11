package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "index2.html", nil)
	if err != nil {
		fmt.Println(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "index3.html", nil)
	if err != nil {
		fmt.Println(err)
	}
}
