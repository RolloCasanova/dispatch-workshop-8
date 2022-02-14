package main

import "fmt"

func main() {
	intValue := 10
	int64Value := int64(20)
	float64Value := float64(30)

	fmt.Println(DescribeVar[int](intValue))
	fmt.Println(DescribeVar[int64](int64Value))
	fmt.Println(DescribeVar[float64](float64Value))

	// // Type argument is not mandatory
	// fmt.Println(DescribeVar(intValue))
	// fmt.Println(DescribeVar(int64Value))
	// fmt.Println(DescribeVar(float64Value))

	// // A type that is not in the list of generic types  wont compile
	// fmt.Println(DescribeVar("40"))
}

func DescribeVar[T int | int64 | float64](v T) string {
	return fmt.Sprintf("Param 'v' is of type %T and its value is: %v", v, v)
}
