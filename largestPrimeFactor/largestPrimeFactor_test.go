// Test Suite for largestPrimeFactor
package largestPrimeFactor

import "testing"

func TestLargestPrimeFactor(t *testing.T) {

	n := Problem3(13195)
	if n != 29 {
		t.Errorf("Wanted %d as larget prime factor of %d, but got %d\n", 29, 13195, n)
	}
}
