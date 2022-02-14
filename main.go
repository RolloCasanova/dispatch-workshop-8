package main

import "fmt"

func main() {
	intValue := 10
	int64Value := int64(20)
	float64Value := float64(30)

	fmt.Println(DescribeInt(intValue))
	fmt.Println(DescribeInt64(int64Value))
	fmt.Println(DescribeFloat64(float64Value))
}

func DescribeInt(v int) string {
	return fmt.Sprintf("Param 'v' is of type int and its value is: %d", v)
}

func DescribeInt64(v int64) string {
	return fmt.Sprintf("Param 'v' is of type int64 and its value is: %d", v)
}

func DescribeFloat64(v float64) string {
	return fmt.Sprintf("Param 'v' is of type float64 and its value is: %f", v)
}
