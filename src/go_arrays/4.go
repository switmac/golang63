package main

import "fmt"

func main() {
	//var x []float64
	//x = append(x,4,5)
	//x := make([]float64, 5)
	//x := make([]float64, 5, 10)
	arr := [5]float64{1,2,3,4,5}
	x := arr[0:5]

	fmt.Println(x)
}
