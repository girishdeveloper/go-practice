package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Paid bool   `json:"paid"`
}

func main() {
	p := Product{Id: 32, Name: "Girish"}
	bI, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bI))
	fmt.Println("---------------------------------------")
	fmt.Println("Marshalling the json values")
	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
