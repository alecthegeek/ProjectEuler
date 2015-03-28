// Problem 4 from Projec Euler
package largestPalindromeProduct

import "strconv"
import "math"

func isPalindrone(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func square(n uint) uint { return n * n }

func Problem4(n uint) (product, f1, f2 uint) {

	// In n digits the largest base 10 number is 10^n - 1
	maxFactor := uint(math.Pow10(int(n))) - 1
	for i := square(maxFactor); i > 0; i-- {
		if isPalindrone(strconv.Itoa(int(i))) {
			for f1 := uint(math.Sqrt(float64(i))); f1 > 0; f1-- {
				if i%f1 == 0 {
					f2 := i / f1
					if f2 <= maxFactor {
						return i, f1, f2
					}
				}
			}
		}
	}

	return 0, 0, 0

}
