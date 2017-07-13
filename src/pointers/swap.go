package main

import "fmt"

func main() {

    x := 5
    y := 6

    swap(&x, &y)

    fmt.Println(x, y)
}

func swap(x, y *int) {
    *x, *y = *y, *x
}
