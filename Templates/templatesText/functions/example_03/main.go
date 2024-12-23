package main

import (
	"fmt"
	"os"
	"text/template"
)

// Встраиваемый шаблон
const Header = `{{block "hello" .}} Text {{end}}`

// блок с ключевым словом define выполняет шаблон hello -> {{ template "hello" .}}
// Замещаем содержимое шаблона "hello"
const Welcome = `
{{define "hello"}}
Здравствуйте и добро пожаловать
{{range .}}{{print .}} {{end}}
{{end}}
`



func main() {
	musketeers := []string{"Атос", "Портос", "Арамис", "Д`Артаньян"}

	helloMsg, err := template.New("start").Parse(Header)
	if err != nil {
		panic(err)
	}

	welcomeMsg, err := template.Must(helloMsg.Clone()).Parse(Welcome)
	if err != nil {
		panic(err)
	}

	if err := helloMsg.Execute(os.Stdout, musketeers); err != nil {
		panic(err)
	}
	fmt.Printf("\n=====\n")

	if err := welcomeMsg.Execute(os.Stdout, musketeers); err != nil {
		panic(err)
	}

}
