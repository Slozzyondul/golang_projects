package variable1

import "fmt"

func VariablesWithoutExpressions() {

	//Variables declared without Initialized ( Without Expression )
	var variable1 int
	var variable2 string
	var variable3 float64
	var variable4 bool

	fmt.Printf("Variable1 Value is : %d\n", variable1)
	fmt.Printf("Variable1 Type is %T\n", variable1)

	fmt.Printf("Variable2 Value is %s\n", variable2)
	fmt.Printf("Variable2 Type is %T\n", variable2)

	fmt.Printf("Variable3 Value is %f\n", variable3)
	fmt.Printf("Variable3 Type is %T\n", variable3)

	//fmt.Printf("Variable4 Value is %yes\n", variable4)
	fmt.Printf("Variable4 Type is %T\n", variable4)

}
