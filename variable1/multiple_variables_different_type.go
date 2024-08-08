package variable1

import "fmt"

func MultipleVariablesDifferentTypes() {

	//s variable declaration by declare multiple variables by differents type in single line
	var1, var2, var3 := 300, "GO Programming is fun with solo", 86.54

	fmt.Printf("The Value of Var 1 : %d\n", var1)
	fmt.Printf("The Type of Var 1 : %T\n", var1)

	fmt.Printf("The Value of Var 2 : %s\n", var2)
	fmt.Printf("The Type of Var 2 : %T\n", var2)

	fmt.Printf("The Value of Var 3 : %f\n", var3)
	fmt.Printf("The Type of Var 3 : %T\n", var3)
}
