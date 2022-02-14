package main

import "fmt"

func main() {
	intValue := 10
	int64Value := int64(20)
	float64Value := float64(30)

	fmt.Println(DescribeVar(intValue))
	fmt.Println(DescribeVar(int64Value))
	fmt.Println(DescribeVar(float64Value))
}

func DescribeVar(v interface{}) string {
	switch v.(type) {
	case int:
		return fmt.Sprintf("Param 'v' is of type int and its value is: %d", v)
	case int64:
		return fmt.Sprintf("Param 'v' is of type int64 and its value is: %d", v)
	case float64:
		return fmt.Sprintf("Param 'v' is of type float64 and its value is: %f", v)
	default: // <- unsupported type
		return fmt.Sprintf("[Unsupported type] Param 'v' is of type %T and its value is: %v", v, v)
	}
}
