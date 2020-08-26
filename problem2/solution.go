package main

import "fmt"

func main() {
	var evenValSum int
	var a = 1
	var b = 2
	var c int
	var i int
	// infinite loop in go
	for {
		// breaking condition, the value should not exceed 4 million
		if c > 4000000 {
			break
		}
		// just to see how many iterations performed
		i++
		// fibonacci logic
		c = a + b
		a = b
		b = c

		// sum even of values
		if c%2 == 0 {
			evenValSum = evenValSum + c
		}
	}
	fmt.Printf("\nTotal iterations: %d\n", i)
	fmt.Printf("Sum of even values: %d\n", evenValSum)
	fmt.Printf("Value of last term: %d\n", c)
}
