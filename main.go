package main

import (
	"fmt"

	"github.com/RolloCasanova/dispatch-workshop-8/stringz"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev := stringz.Reverse(input)
	doubleRev := stringz.Reverse(rev)

	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}
