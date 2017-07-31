package calculator

import "fmt"

func init() {
    fmt.Println("calculator initialised")
}

func Add (x, y float64) float64 {
    return x + y;
}

func Multiply (x, y float64) float64 {
    return x * y;
}
