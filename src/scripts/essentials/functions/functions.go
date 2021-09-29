package main

import "fmt"

func add(a float64, b float64) float64 {
	return a + b
}

func main() {

	a := 2
	b := 2
	sum := add(2, 2)

	// `int` is required because the `add` function
	// has explicitly typed `sum` as `float64`
	// but the formatter `%d` expects an `int` type
	fmt.Printf("%d plus %d is %d", a, b, int(sum))
}
