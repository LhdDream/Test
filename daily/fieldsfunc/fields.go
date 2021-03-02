package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	f := func(c rune) bool {
		return unicode.IsNumber(c)
	}

	fmt.Println("Fields are %q", strings.FieldsFunc("ABC123PQ456XZ789",f))
}
