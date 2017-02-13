package main

import (
	"html/template"
	"os"
)

type Person struct {
	Username string
}

func main() {
	t := template.New("Template Example")
	t, _ = t.Parse(" Hello, World. My name is {{.Username}}")
	p := Person{Username: "Darwin"}

	t.Execute(os.Stdout, p)
}
