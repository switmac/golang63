package main

import "fmt"

var z = 1

func main()  {
  //integer
  fmt.Println("1 + 1 =", 1 + 1)

  //floating
  fmt.Println("1 + 1 =", 1.0 + 1.0)

  var x = "Hello"
  y:= "World"

  fmt.Println(x + " "+ y)


}

func sum()  {
  fmt.Println(z + 3)
}
