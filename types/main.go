package main

import "fmt"

func main() {
	integer()
	float()
	string()
}

func integer() {
	fmt.Println("1+1=", 1+1)
}
func float() {
	fmt.Println("1+1=", 1.0+1.0)
}
func string() {
	fmt.Println(len("Helllo world"))
	fmt.Println("hello world"[1])
	fmt.Println("hello" + " world")
}
