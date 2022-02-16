package main

import (
	"constraints"
	"fmt"
)

func main() {
	x1, y1 := 1, 2
	fmt.Println(isLower(x1, y1))

	x2, y2 := -5.0, -10.0
	fmt.Println(isLower(x2, y2))

	x3, y3 := "a", "b"
	fmt.Println(isLower(x3, y3))
}

// isLower returns true if x is lower than y, otherwise returns false
func isLower[T constraints.Ordered](x, y T) bool {
	return x < y
}
