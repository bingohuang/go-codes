package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.TrimRight("abcdcbab", "abc"))
	fmt.Println(strings.TrimRight("abcdcbab", "a"))
	fmt.Println(strings.TrimRight("abcdcbab", "b"))
	fmt.Println(strings.TrimRight("¡¡¡Hello, Gophers!!!", "!¡"))
	fmt.Println(strings.TrimRight("abcbab", "ab"))
}
