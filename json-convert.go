package main

import (
	"encoding/json"
	"fmt"
)

type CoolObject struct {
	Name         string `json:"Enam"`
	Age          int    `json:"Ega"`
	DateOfBirth  string `json:"DOB,omitempty"`
	OutdatedInfo string `json:"-"`
}

func main() {
	data := []byte(`{ "Enam":"Niwrad", "Ega":32 }`)

	var obj CoolObject
	json.Unmarshal(data, &obj)

	fmt.Println(obj)
}
