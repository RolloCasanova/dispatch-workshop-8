package main

import "fmt"

// Signed is a constraint whose type set is any signed integer type.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a constraint whose type set is any unsigned integer type.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Float is a constraint whose type set is any floating point type.
type Float interface {
	~float32 | ~float64
}

// Ordered is a constraint whose type set is any ordered type.
// That is, any type that supports the < operator.
type Ordered interface {
	Signed | Unsigned | Float | ~string
}

func main() {
	x1, y1 := 1, 2
	fmt.Println(isLower(x1, y1))

	x2, y2 := -5.0, -10.0
	fmt.Println(isLower(x2, y2))

	x3, y3 := "a", "b"
	fmt.Println(isLower(x3, y3))
}

// isLower returns true if x is lower than y, otherwise returns false
func isLower[T Ordered](x, y T) bool {
	return x < y
}
