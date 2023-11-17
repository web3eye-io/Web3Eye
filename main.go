package main

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {

	things := []string{"foo", "bar", "baz"}
	fmt.Println(slices.ContainsFunc[string](things, func(s string) bool {
		return strings.Contains(s, "fo")
	})) // true

}
