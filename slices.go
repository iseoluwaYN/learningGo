package main
import "fmt"

//creating a slice
var a []float64

//second slice
func second(){
	b := make([]float64, 5)
	fmt.Println(b)
}

//third slice
func third(){
	c := make([]float64, 5, 11)
	fmt.Println(c)
}
func sliceAlternative(){
	arr := []float64{1,2,3,4,5}
	d:= arr[0:5]
	fmt.Println(d)
}

func appending(){
	e :=[]float64{2,3,5,4,6,6}
	f := append(e,3,3,5,6,8,5,6,3,5,2)
	fmt.Println(f)
}

func copying(){
	slice1 := []int{1,2,3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

func leastNumber(){
	x := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}
	var smallest =0
	for i :=0;i < len(x); i++{
		var initialNumber = i
		nextNumber := x[i+1]
		if initialNumber < nextNumber {
			smallest = initialNumber
		}
	}
	fmt.Println(smallest)
}