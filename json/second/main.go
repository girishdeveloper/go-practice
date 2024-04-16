package main

import (
	"encoding/json"
	"fmt"
)

type MyJson struct {
	Cat `json:"cat"`
}

type Cat struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   uint64 `json:"age"`
}

func main() {
	myjson := []byte(`{"cat": {"name": "Khoye", "breed": "Himalayan", "age": 42}}`)
	c := MyJson{}
	err := json.Unmarshal(myjson, &c)
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Cat.Name)
	fmt.Println(c.Cat.Breed)
	fmt.Println(c.Cat.Age)
}
