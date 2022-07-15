package main

//
//import (
//	"fmt"
//)
//var s = []float64{98,93,77}
//func main() {
//	first()
//	average(s)
//	add()
//	defer second()
//	firstt()
//
//	defer func() {
//		str := recover()
//		fmt.Println(str)
//	}()
//	panic("PANIC")
//}
//
//func first(){
//	xs := []float64{98,93,77,82,83}
//	total := 0.0
//	for _, v := range xs {
//		total += v
//	}
//	fmt.Println(total / float64(len(xs)))
//}
//
//func average(xs []float64) float64 {
//	total := 0.0
//	for _, v := range xs {
//		total += v
//	}
//	return total / float64(len(xs))
//}
//
//func add(args ... int ) int {
//	total:=0
//	for _, v:= range args{
//		total += v
//	}
//	return  total
//}
//
//func firstt() {
//	fmt.Println("1st")
//}
//func second() {
//	fmt.Println("2nd")
//}
