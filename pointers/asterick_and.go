package main

import "fmt"

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1
}
func zero(xPtr *int) {
	*xPtr = 0
}
func one(xPtr *int) {
	*xPtr = 1
}