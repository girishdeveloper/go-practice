package main

import (
	"fmt"
)

func main() {
	var emailSent, emailToSend int = 0, 3
	for emailSent < emailToSend {
		fmt.Println("Email sent", emailSent+1, "times")
		emailSent++
	}
}
