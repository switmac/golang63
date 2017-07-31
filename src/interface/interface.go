package main

import "fmt"

type Person struct {
    name string
    age int
    skill string
}

type Robot struct {
    name string
    age int
    model string
    skill string
}

type Skill interface {
    Show() string
}

func (p Person) Show() string {
    return p.skill
}

func (r Robot) Show() string {
    return r.skill
}

func main() {

    arnel := Person { "arnel", 16, "jogging" }

    dave := Robot { "kent", 16, "3210", "flying" }

    fmt.Println("im a person my skill is ", arnel.Show())
    fmt.Println("im a robot my skill is ", dave.Show())

    fmt.Println(arnel)

}
