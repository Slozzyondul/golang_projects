package variable1

import "fmt"

func Variables1() {
	// Variable declared and initialized without the explicit type
	var variable1 = 100
	var variable2 = "GO Programming Language"
	var variable3 = 750.345

	// Display and print variables
	fmt.Printf("variable1 Value is : %d\n", variable1)
	fmt.Printf("Variable1 Type is %T\n", variable1)

	fmt.Printf("Variable2 Value is %s\n", variable2)
	fmt.Printf("Variable2 Type is %T\n", variable2)

	fmt.Printf("Variable3 Value is %f\n", variable3)
	fmt.Printf("Variable3 Type is %T\n", variable3)
}
