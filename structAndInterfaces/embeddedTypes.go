package main

import "fmt"

//import "fmt"
var human Human

func main() {
 	words := "Iseoluwa is dope"
 	fmt.Println(
 	human.speak(words))

// 	inheritance
	tobi := new(Student)
	tobi.speak(words)
}

type Human struct {
	Name string
	Age int16
	Gender string
}

func (human Human) speak(words string) string {
	return words
}

// inheritance
type Student struct {
	Human
	Class string
}


