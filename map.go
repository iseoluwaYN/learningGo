package main
import "fmt"

func maps() {
	var x =make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])
	//to delete
	delete(x, "key")
}

//first 20 elements
func elements(){
	elements := make(map[string]string)
	elements["H"]= "hydrogen"
	elements["He"]= "helium"
	elements["Li"]= "Lithium"
	fmt.Println(elements["Li"])
}

//modified version
func mapModified(){
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