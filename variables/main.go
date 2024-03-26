package main

import "fmt"

func main() {
	// var myName string = "Umar"
	myName := "Umar"
	// gotcha in go. if variable is declared and not used. go will not compile it.
	myNumber := 16
	myFloat := 2.0
	fmt.Printf("Hello my name is %s\nmy number is %d\nand my float is %f", myName, myNumber, myFloat)
}
