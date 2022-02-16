package stringz

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return // or you can call t.Skip() to stop the execution of that fizz input
		}
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

// to test a specific fuzz funciton by using seed corpus:
// go1.18beta2 test -v ./stringz --run=FuzzReverse

// to fuzzy test a specific fuzz function:
// go1.18beta2 test ./stringz --fuzz=Fuzz --run=FuzzReverse

// fuzz test wont stop until it finds an error, or a timeout:
// define a timeout for the test, or use Ctrl + C to stop test execution
// go1.18beta2 test ./stringz --fuzz=Fuzz --run=FuzzReverse --fuzztime=10s
