package main

import "fmt"

func main() {
	/*var x map[string]int
	x["key"] = 10
	fmt.Println(x)*/

	/*x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x)*/

	x := make(map[int]int)
	x[1] = 10
	fmt.Println(x[1])
	delete(x, 1)
	fmt.Println(x)
	fmt.Println(len(x))
}
