package datatypes

import "fmt"

func FloatType() {

	//define floating-point numbers

	a := 5635.7868798657587
	b := 46757865.2655

	c := b - a

	fmt.Printf("A = %f\n", a)
	fmt.Printf("B = %f\n", b)

	fmt.Printf("Type of C : %T\n", c)

}
