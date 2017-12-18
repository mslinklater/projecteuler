/*
Problem 1 - Multiples of 3 and 5

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get
3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multples of 3 or 5 below 1000.
*/

package main

import "fmt"

func main() {
	fmt.Println("ProjectEuler Problem 1")

	var sum int
	for num := 1; num < 1000; num++ {
		if (num%3 == 0) || (num%5 == 0) {
			sum += num
		}
	}
	fmt.Println("Answer:", sum)
}
