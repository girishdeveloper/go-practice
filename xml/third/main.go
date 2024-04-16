package main

import (
	"encoding/xml"
	"fmt"
)

type Price struct {
	Text     string `xml:",chardata"`
	Currency string `xml:"currency,attr"`
}

type Product struct {
	Name       string  `xml:"name"`
	Price      Price   `xml:"price"`
	Commission float64 `xml:,cdata`
	Comment    string  `xml:",comment"`
}

func main() {
	c := Product{}
	c.Price = Price{Text: "123.94", Currency: "EUR"}
	c.Name = "Test Data"
	c.Comment = "This is my comment"
	c.Commission = 9.04
	bI, err := xml.MarshalIndent(c, "", "	")
	if err != nil {
		panic(err)
	}
	output := xml.Header + string(bI)
	fmt.Println("with attributes", output)
}
