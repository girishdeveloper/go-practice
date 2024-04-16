package main

import (
	"encoding/xml"
	"fmt"
)

type MyXml struct {
	Cat `xml:cat`
}

type Cat struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}

func main() {
	p := []byte(`<cat>
    <name>Ti</name>
    <age>23</age>
</cat>`)
	c := MyXml{}
	err := xml.Unmarshal(p, &c)
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Cat.Name)
	fmt.Println(c.Cat.Age)
}
