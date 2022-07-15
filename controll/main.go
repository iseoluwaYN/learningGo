package main

import "fmt"

func main() {
	//moreLoops()
	//oddNumbers()
	primeNumbers()
}

func loops() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
}

func moreLoops() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func oddNumbers() {
	fmt.Println("Even and odd numbers from 1- 20")
	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i, " is an even number.")
		} else {
			fmt.Println(i, " is an odd number.")
		}
	}
}

func primeNumbers() {
	fmt.Print(2)
	for i := 0; i <= 20; i++ {
		if i%2 == 0 && i%1 == 0 {

		}
	}
}
