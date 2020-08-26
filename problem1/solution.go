package main

import "fmt"

func main() {
	// initialize sum var to hold the result
	var sum = 0
	// we will start the iteration from 3 as this is the first
	// natural number that is multiple of 3
	for n := 3; n < 1000; n++ {
		if (n%3 == 0) || (n%5 == 0) {
			fmt.Printf("\n%d + %d", sum, n)
			sum = sum + n
		}
	}

	// print the result
	fmt.Println("the sum of all the multiples of 3 or 5 below 1000 are: ", sum)
}
