package datatypes

import "fmt"

func ComplexCompileType() {

	var a = 2.353

	var b = 3.76778

	var c = complex(a, b)

	fmt.Println(c)
	fmt.Printf("Type of C : %T\n", c)
}
