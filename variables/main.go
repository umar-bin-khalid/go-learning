package main

import "fmt"

func main() {
	// var myName string = "Umar"
	myName := "Umar"
	// gotcha in go. if variable is declared and not used. go will not compile it.

	fmt.Printf("Hello my name is %s\n", myName)
}
