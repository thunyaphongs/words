package words

import (
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s) // HL
	r := map[string]int{}
	for _, w := range words {
		r[w] = r[w] + 1 // HL
	}
	return r
}
