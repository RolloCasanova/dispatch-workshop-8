package main

import "fmt"

// type myCustomInt int
// type myOtherCustomInt int

// type customInts interface {
// 	myCustomInt | myOtherCustomInt
// }

func main() {
	intValue := 10
	int64Value := int64(20)
	float64Value := float64(30)
	stringValue := "40"
	arrayValue := []rune{'a', 'b', 'c'}
	boolValue := true
	mapValue := map[string]string{"k1": "v1"}
	structValue := struct{}{}

	fmt.Println(DescribeVar(intValue))
	fmt.Println(DescribeVar(int64Value))
	fmt.Println(DescribeVar(float64Value))
	fmt.Println(DescribeVar(stringValue))
	fmt.Println(DescribeVar(arrayValue))
	fmt.Println(DescribeVar(boolValue))
	fmt.Println(DescribeVar(mapValue))
	fmt.Println(DescribeVar(structValue))

	// // for tilde operator
	// var customInt myCustomInt = 100
	// var otherCustomInt myOtherCustomInt = 200

	// fmt.Println(DescribeVar(customInt))
	// fmt.Println(DescribeVar(otherCustomInt))
}

// Hardcode the types of the variable
func DescribeVar[T int | int64 | float64 | string | []rune | bool | map[string]string | struct{}](v T) string {
	return fmt.Sprintf("Param 'v' is of type %T and its value is: %v", v, v)
}

// // Interface approach
// type myValues interface {
// 	int | int64 | float64 |
// 	string | []rune | bool |
// 	map[string]string | struct{}
// }

// func DescribeVar[T myValues](v T) string {
// 	return fmt.Sprintf("Param 'v' is of type %T and its value is: %v", v, v)
// }

// // Any (interface{}) apprach
// func DescribeVar[T any](v T) string {
// 	return fmt.Sprintf("Param 'v' is of type %T and its value is: %v", v, v)
// }

// func DescribeVar[T interface{}](v T) string {
// 	return fmt.Sprintf("Param 'v' is of type %T and its value is: %v", v, v)
// }

// // Custom types (tilde operator)
// func DescribeVar[T int | myCustomInt | myOtherCustomInt](v T) string {
// 	return fmt.Sprintf("Param 'v' is of type %T and its value is: %v", v, v)
// }

// func DescribeVar[T ~int](v T) string {
// 	return fmt.Sprintf("Param 'v' is of type %T and its value is: %v", v, v)
// }
