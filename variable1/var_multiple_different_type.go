package variable1

import "fmt"

func MultipleVariables() {

	//declare multiple variables for different type in the single diclaration
	var character1, character2, character3 = 2, "Go is interesting", 67.56

	fmt.Printf("character1 Value : %d\n", character1)
	fmt.Printf("character1 Type : %T\n", character1)

	fmt.Printf("character2 Value : %s\n", character2)
	fmt.Printf("character2 Type : %T\n", character2)

	fmt.Printf("character3 Value : %f\n", character3)
	fmt.Printf("character3 Type : %T", character3)

}
