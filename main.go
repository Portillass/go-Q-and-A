/*Q1: Write a simple “Hello World” program in Go.*/


package main // Difines the main package - entry poiny of evert Go Program

import (
	"fmt"
	"example.com/goQandA/Q2"
	"example.com/goQandA/Q3"
)
func main() {
	 	fmt.Println("Hello, World!") // Print Hello, World! to the console
	Q2.Add()
	Q3.Subtract()
}

