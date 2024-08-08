package variable1

import "fmt"

var golobalVariable int = 100

func LocalandGlobalVariables() {
	//local variable scope starts

	var localVariable int = 200

	fmt.Printf("Global Variable is %d\n", golobalVariable)

	fmt.Printf("Local Variable : %d\n", localVariable)

	display()

	// local variable scope ends
}

func display() {

	fmt.Printf("The Value of Global variable is %d\n", golobalVariable)
}
