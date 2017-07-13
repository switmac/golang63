package main

import "fmt" 

func a() int {
    x := 3
    return x
}

// named return var
func b() (y int) {
    y = 99
    return
}

// returning multiple values
func c() (int, int, int) {
    return 99, 45, 33
}
func c2() (string, int, float64) {
    return "string", 44, 99.99
}
func c3() []float64 {
    x := []float64{11.11, 22.22, 33.33}
    return x
}

// variadic funcs
func d(args ...int) int {
    total := 0
    for _, x := range args {
        total += x
    }
    return total
}

func main() {
    fmt.Println(a())

    fmt.Println(b())

    fmt.Println(c())
    fmt.Println(c2())
    fmt.Println(c3())

    fmt.Println(d(1,2,3,4,5,6))
    fmt.Println(d(99,0,123))
    fmt.Println(d())
}
