package main

import "fmt"

func main() {
	//array()
	//specialLoop()
	//maps()
	//elements()
	//mapModified()

	//var x [5]float64
	//x := make([]float64, 5, 10)

	arr := []float64{1, 2, 3, 4, 5}
	x := arr[2:len(arr)]
	fmt.Print(x)
}

var x = [5]float64{98, 93, 77, 82, 83}
var total float64 = 0

func array() {
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
}
func specialLoop() {
	//var x [5] float64
	//x[0] = 98; x[1] = 93; x[2] = 77; x[3] = 82; x[4] = 83
	for _, value := range x {
		total += value
		fmt.Println(total / float64(len(x)))
	}
}
func maps() {
	var x = make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])
	//to delete
	delete(x, "key")
}

//first 20 elements
func elements() {
	elements := make(map[string]string)
	elements["H"] = "hydrogen"
	elements["He"] = "helium"
	elements["Li"] = "Lithium"
	fmt.Println(elements["Li"])
}

//modified version
func mapModified() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
