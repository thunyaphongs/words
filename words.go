package words

import (
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	fmt.Println("Au")
	words := strings.Fields(s) // HL
	r := map[string]int{}
	for _, w := range words {
		r[w] = r[w] + 1 // HL
	}
	return r
}
