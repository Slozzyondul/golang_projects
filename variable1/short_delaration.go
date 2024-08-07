package variable1

import "fmt"

func ShortDeclaration() {

	//type of the variable is determined by the type of expression
	variable1 := 100
	variable2 := "GO Programming is so cool for solo"
	variable3 := 135.94

	fmt.Printf("Variable1 Value : %d\n", variable1)
	fmt.Printf("Variable1 Type : %T\n", variable1)

	fmt.Printf("Variable2 Value : %s\n", variable2)
	fmt.Printf("Variable2 Type : %T\n", variable2)

	fmt.Printf("Variable3 Value : %f\n", variable3)
	fmt.Printf("Variable3 Type : %T\n", variable3)
}
