// Return and word:wordcount map for a given string

package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	count_map := make(map[string]int)
	for _, val := range words {
		_, exists := count_map[val]
		if exists {
			count_map[val] += 1
		} else {
			count_map[val] = 1
		}
	}
	return count_map
}

func main() {
	fmt.Println(WordCount("I am groot"))
	fmt.Println(WordCount("I ate a donut. Then I ate another donut."))
	fmt.Println(WordCount("A man a plan a canal panama."))
}
