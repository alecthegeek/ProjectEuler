// Brief one line overview of application
package main

import "fmt"

import (
	"github.com/alecthegeek/ProjectEuler/evenFibonacciNumbers"
	"github.com/alecthegeek/ProjectEuler/largestPalindromeProduct"
	"github.com/alecthegeek/ProjectEuler/largestPrimeFactor"
	"github.com/alecthegeek/ProjectEuler/multiplesOf3and5"
)

func main() {

	fmt.Printf("Running %s: Problem %d  Result is %d\n",
		"Multiples of 3 and 5", 1, multiplesOf3and5.Problem1(1000))

	fmt.Printf("Running %s: Problem %d  Result is %d\n",
		"Even Fibonacci Numbers", 2, evenFibonacciNumbers.Problem2(4000000))

	fmt.Printf("Running %s: Problem %d  Result is %d\n",
		"Largest prime factor", 3, largestPrimeFactor.Problem3(600851475143))

	p, f1, f2 := largestPalindromeProduct.Problem4(3)
	fmt.Printf("Running %s: Problem %d  Result is %v = %v x %v\n",
		"Largest Palidronic product", 4, p, f1, f2)

}
