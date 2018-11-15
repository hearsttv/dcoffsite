package main

import (
	"fmt"
)

func helloDC(done chan bool) {

	fmt.Println("=^..^=")
	done <- true
}
func main() {
	done := make(chan bool)
	go helloDC(done)
	<-done
	fmt.Println("main function")
}
