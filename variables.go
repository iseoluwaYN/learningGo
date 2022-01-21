package main
import "fmt"

func variable() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}

func convertFromFahrenheits(){
	fmt.Println("Enter degree in Fahrenheits: ")
	var F float64
	fmt.Scanf("%f", &F)
	C := (F - 32) * (5/9)
	fmt.Println(C)
}

func feetToMeters(){
	fmt.Println("Enter height in feet")
	var feet float64
	fmt.Scanf("%f", &feet)
	meters := feet * 0.3048
	fmt.Println("You are ",meters, " meter tall")
}