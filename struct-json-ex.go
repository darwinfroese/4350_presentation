package main

import "encoding/json"

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "darwin", Age: 23}

	json.Marshal(person)
}
