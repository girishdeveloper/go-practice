package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Employee struct {
	Name     string
	Age      int
	Position string
}

func IterateMap(employees map[string]Employee) {
	for key, value := range employees {
		fmt.Printf("%v, %v, %v, %v\n", key, value.Name, value.Age, value.Position)
	}
}

func main() {
	data := make(map[string]Employee)
	file, err := os.Open("./employees.csv")
	if err != nil {
		log.Fatalf("Error at file opening: %s", err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	i := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if i == 0 {
			fmt.Println(record[0], record[1], record[2], record[3])
			fmt.Println("-----------------------------")
			i++
		} else {
			employeeAge, err := strconv.Atoi(record[2])
			if err == nil {
				data[record[0]] = Employee{
					Name:     record[1],
					Age:      employeeAge,
					Position: record[3],
				}
			} else {
				log.Fatal(err)
			}
		}
	} // end of for
	IterateMap(data)
	//Retrieve a record from the map
	value, ok := data["V0003"]
	if ok {
		fmt.Println(value, ok)
	} else {
		log.Fatalf("No data for V0003 %s", ok)
	}
	//Delete a record from the map
	delete(data, "V0002")
	IterateMap(data)
}
