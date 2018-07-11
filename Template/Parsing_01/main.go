package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		fmt.Println(err)
	}
	fl, err := os.Create("index2.html")
	if err != nil {
		fmt.Println(err)
	}
	err = tpl.Execute(fl, nil)
	if err != nil {
		fmt.Println(err)
	}
}
