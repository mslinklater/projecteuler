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
