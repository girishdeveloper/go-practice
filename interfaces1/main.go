package main

import (
	"dbo/driver"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Interface for database connection")
	log.Println("Start")
	ConDetail := &driver.ConnectionDetail{Hostname: "server1", Username: "girish", Password: "egi_123", Database: "dataOne"}
	log.Println(ConDetail)
	Conn, err := ConDetail.Open()
	if err != nil {
		log.Println("Error while opening connection")
		return
	}
	log.Println("connection string is", Conn)
	result, err := ConDetail.Execute(ConDetail.ConId, "SELECT * FROM tbl_data")
	if err != nil {
		log.Println("Error while executing a SQL statement")
	}
	log.Println("Query results")
	log.Println(result)
	ConDetail.Close(ConDetail.ConId)
	log.Println("STOP")
}
