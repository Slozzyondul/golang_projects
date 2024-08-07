package variable1

import "fmt"

func GroupedBlocks() {

	//define a groups of variables
	var (
		firstname = "solo"
		age       = 27
		number2   = +254792352745
		married   = true
	)

	fmt.Println("Firstname : ", firstname)
	fmt.Println("Number1 : ", age)
	fmt.Println("Number2 : ", number2)
	fmt.Println("Married : ", married)
}
