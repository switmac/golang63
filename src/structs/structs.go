package main

import "fmt"

type Person struct {
    name string
    age int
}

func main() {

    arnel := Person { name: "arnel", age: 18 }

    dave := Person { "name", 19 }

    marie := struct {
        name string
        age int
    } { "marie", 20  }

    fmt.Println(arnel)
    fmt.Println(dave)
    fmt.Println(marie)
}
