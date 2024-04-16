package main

import (
	"fmt"
	"os"
)

func main() {
	myvar := os.Getenv("MYVAR")
	fmt.Printf("MYVAR = %s\n", myvar)
	myvar2, found := os.LookupEnv("DB_PORT")
	if !found {
		fmt.Println("value not set")
		myvar2 = "8081"
		os.Setenv("DB_PORT", myvar2)
	}
	fmt.Println(os.Environ())
}
