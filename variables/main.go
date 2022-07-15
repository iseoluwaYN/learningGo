package main

import "fmt"

func main() {
	var x string
	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)

	x = "darling"

	name := "Max"
	fmt.Println("My dog's name is", name)

	fmt.Print("Enter a two digit number ")
	var input float64
	fmt.Scanf("%f", input)
	multiple := input * 2
	fmt.Print(multiple)
}
