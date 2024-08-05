package variable1

import "fmt"

func SameTypeVariables() {

	//define multiple variables of the same type are declared and initialized in the single line
	var variable1, variable2, variable3 int = 2122, 4561, 6723423
	var varibale4, variable5, variable6 string = "solo", "sikoa", "tribal chief"

	fmt.Printf("Variable1 Value : %d\n", variable1)
	fmt.Printf("Variable2 Value : %d\n", variable2)
	fmt.Printf("Variable3 Value : %d\n", variable3)
	fmt.Printf("%s\n", varibale4)
	fmt.Printf("%s\n", variable5)
	fmt.Printf("%s\n", variable6)

}
