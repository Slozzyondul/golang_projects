package datatypes

import "fmt"

func StringType() {

	//declare string data type
	str := "Go Programming Language is very interesting to know for the future"

	fmt.Printf("The Value of str : %s\n", str)

	fmt.Printf("The Type of str : %T\n", str)

	fmt.Printf("The Length of str : %d\n", len(str))
}
