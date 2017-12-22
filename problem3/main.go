/*
Problem 3 - Largest Prime Factor

The prime factors of 13195 are 5, 7, 13 and 29
What is the largest prime factor of the number 600851475143
*/
package main

import (
	"fmt"
	"math"
	"time"
)

func isPrime(num int64) bool {
	// so the idea here is to iterate from 2 to sqrt(num) and see if the number is a factor
	limit := int64(math.Sqrt(float64(num)))

	var testNum int64
	testNum = 2
	for testNum <= limit {
		if num%testNum == 0 {
			return false
		}
		testNum++
	}
	fmt.Printf("%d is prime\n", num)
	return true
}

func main() {
	fmt.Println("ProjectEuler Problem 3")

	var theNumber int64
	theNumber = 600851475143

	var i int64
	t0 := time.Now()
	for i = 2; i < theNumber; i++ {
		if theNumber%i == 0 {
			testNum := theNumber / i
			if isPrime(testNum) {
				t1 := time.Now()
				fmt.Printf("Found it:%d (%fs)", testNum, t1.Sub(t0).Seconds())
				return
			}
		}
	}
}
