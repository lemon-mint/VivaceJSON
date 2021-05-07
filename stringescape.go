package main

import (
	"fmt"
	"strconv"
)

func gencalc(a []*field) int {
	c := 0
	return c
}

func calcString(a []byte) int {
	c := 0

	for i := range a {
		switch {
		case a[i] == '"':
			c += 2 // => \"
		case a[i] == '\\':
			c += 2 // => \\
		case a[i] == '\n':
			c += 2 // => \n
		case a[i] == '\t':
			c += 2 // => \t
		case a[i] == '\r':
			c += 2 // => \r
		case a[i] < 32:
			c += 6 // => \u00XX
		default:
			c += 1
		}
	}

	return c
}

func calcInt(a int) int {
	return len(strconv.Itoa(a))
}

func calcBool(a bool) int {
	if a {
		return 4
	}
	return 5
}

func calcFloat(a float64) int {
	return len(fmt.Sprintf("%f", a))
}
