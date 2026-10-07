package main

import "fmt"

// join concatenates parts with commas.
func join(parts []string) string {
	out := ""
	for i, p := range parts {
		if i > 0 {
			out += ","
		}
		out += p
	}
	return out
}

func main() { fmt.Println(join([]string{"a", "b", "c"})) }
