package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type Content struct {
	Name     string
	Age      int
	Position string
}

func main() {
	var EmployeeDetail = make(map[string]Content)
	file, err := os.Create("./example.txt")
	if err != nil {
		log.Fatalf("file could not be created: %s", err)
	}
	defer file.Close()
	w := csv.NewWriter(file)
	EmployeeDetail["V0001"] = Content{
		Name:     "Girish",
		Age:      42,
		Position: "Vice President",
	}
	EmployeeDetail["V0002"] = Content{
		Name:     "Koresh",
		Age:      40,
		Position: "Managing Director",
	}
	employeeSlice := [][]string{
		{"V0001", EmployeeDetail["V0001"].Name, strconv.Itoa(EmployeeDetail["V0001"].Age), EmployeeDetail["V0001"].Position},
		{"V0002", EmployeeDetail["V0002"].Name, strconv.Itoa(EmployeeDetail["V0002"].Age), EmployeeDetail["V0002"].Position},
	}
	for _, row := range employeeSlice {
		_ = w.Write(row)
		/*if err != nil {
			log.Fatalf("writing error: %s", err)
		}*/
	}
	w.Flush()
}
