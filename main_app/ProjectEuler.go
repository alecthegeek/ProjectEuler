// Brief one line overview of application
package main

import "fmt"

import (
	"github.com/alecthegeek/ProjectEuler/multiplesOf3and5"
)

func main() {

	fmt.Printf("Running %s: Problem %d  Result is %d\n",
		"Multiples of 3 and 5", 1, multiplesOf3and5.Problem1(1000))
}
