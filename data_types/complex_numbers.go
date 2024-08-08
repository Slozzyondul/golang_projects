package datatypes

import "fmt"

func ComplexNumbersType() {

	var a = 2453 + 985i
	var b = 12 - 43i

	var result1 = a + b

	var result2 = b - a

	var result3 = a * b

	var result4 = a / b

	fmt.Println(result1, result2, result3, result4)
}
