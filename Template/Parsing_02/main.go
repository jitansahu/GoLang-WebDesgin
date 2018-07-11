package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("index.gohtml", "index2.html")
	if err != nil {
		fmt.Println(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "index2.html", nil)
	if err != nil {
		fmt.Println(err)
	}
}
