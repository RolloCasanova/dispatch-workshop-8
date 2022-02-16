package stringz

import "log"

func Reverse(s string) string {
	log.Printf("input: %q\n", s)
	r := []rune(s)
	log.Printf("runes: %q\n", r)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
