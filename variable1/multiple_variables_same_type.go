package variable1

import "fmt"

func MultipleVariablesSameType() {

	//short variable declaration by multiple variables of same type in single declaration
	var1, var2, var3 := 100, 200, 300

	fmt.Printf("The Value of Var 1 : %d\n", var1)
	fmt.Printf("The Type of Var 1 : %T\n", var1)

	fmt.Printf("The Value of Var 2 : %d\n", var2)
	fmt.Printf("The Type of Var 2 : %T\n", var2)

	fmt.Printf("The Value of Var 3 : %d\n", var3)
	fmt.Printf("The Type of Var 3 : %T\n", var3)
}
