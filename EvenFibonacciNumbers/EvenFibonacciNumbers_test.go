// Breif one line overview of package
package EvenFibonacciNumbers

import "testing"
import "fmt"

var data = []struct{ n, r uint }{
	{0, 0},
	{1, 0},
	{2, 0},
	{3, 2},
	{5, 2},
	{89, 44},
}

func TestProblem2(t *testing.T) {
	for _, d := range data {
		if got := Problem2(d.n); got != d.r {
			t.Errorf(fmt.Sprintf("Problem1 of %d returned %d, expected %d\n", d.n, got, d.r))
		}
	}
}
