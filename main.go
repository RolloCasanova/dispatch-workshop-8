package main

import (
	"fmt"

	"github.com/RolloCasanova/dispatch-workshop-8/stringz"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, revErr := stringz.Reverse(input)
	doubleRev, doubleRevErr := stringz.Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}
