package main

import (
	"encoding/xml"
	"fmt"
)

type Cat struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

func main() {
	p := Cat{Name: "Meow", Age: 8}
	bI, err := xml.MarshalIndent(p, "", "	")
	if err != nil {
		panic(err)
	}
	xmlWithHeader := xml.Header + string(bI)
	fmt.Println(xmlWithHeader)
}
