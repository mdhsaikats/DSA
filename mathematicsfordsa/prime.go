package main

import (
	"fmt"
	"math"
)

func primeNumber(n int) {
	if n <= 1 {
		fmt.Println("It is not a prime number")
		return
	}

	num := int(math.Sqrt(float64(n)))

	for i := 2; i <= num; i++ {
		if n%i == 0 {
			fmt.Println("It is not a prime number")
			return
		}
	}

	fmt.Println("It is a prime number")
}

func main() {
	primeNumber(2)
	primeNumber(9)
	primeNumber(17)
}