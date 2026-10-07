package main

import (
	"fmt"
	"strings"
)

// join concatenates parts with commas.
func join(parts []string) string {
	return strings.Join(parts, ",")
}

func main() { fmt.Println(join([]string{"a", "b", "c"})) }
