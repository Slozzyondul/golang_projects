package variable1

import "fmt"

var variable1 int = 500

func LocalGlobalVariable() {

	var variable1 int = 200

	fmt.Printf("Global Variable is %d\n", variable1)

}
