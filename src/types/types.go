package main //package declaration this line is required

import "fmt" //sample library package usage, format

 /*
 * Multi line comment
 */

var z int = 5 //global variable declaration

func main()  {
  //method call
  sample_variable_declaration()
  sample_global_var()
  sample_input()
  sample_if()
  sample_for()
  sample_switch()
}

func sample_variable_declaration() {
	var z float64 = 45.6

	//integer
	fmt.Println("1 + z =", z + 1)

	//floating
	fmt.Println("1 + 1 =", 1.0 + 1.0)

	var x = "Hello"
	y:= "World"

	is_bool:= true

	fmt.Println(x + " "+ y)
	fmt.Println(is_bool)
}

func sample_global_var()  {
	fmt.Println("sample access global variable")
	
  	fmt.Println("sum ", z + 3)
}

func sample_input() {
	var (
		menuend float64
		subtrahend float64
	)
	
	fmt.Println("sample input")
	
	fmt.Scanf("%f",   &menuend)
	fmt.Scanf("%f",   &subtrahend)

	diff := menuend - subtrahend
	fmt.Println(menuend, " - ", subtrahend, " = ", diff)  
}

func sample_if() {
	var value int
	
	fmt.Println("sample if")
	
	fmt.Scanf("%f",   &value)

	if (value % 2) == 0 {
		fmt.Println(value, "is even")		
	} else {
		fmt.Println(value, "is odd")
	}		
}


func sample_for() {
	var loopcount int
	
	fmt.Println("sample for")
	
	fmt.Scanf("%f", &loopcount)

	for i := 1; i <= loopcount; i++ {
		fmt.Println(i)		
	}

	j:= 1
	for j <= 5 {
		fmt.Println(j)
		j = j + 1 			
	}	
}

func sample_switch() {
	var input int
	
	fmt.Println("sample switch")
	
	fmt.Scanf("%f", &input)

	switch input {
		case 1:  	fmt.Println("This is number ", input)
		case 2: 	fmt.Println("This is number ", input)
		default: 	fmt.Println("Unsupported input") 
	}
}