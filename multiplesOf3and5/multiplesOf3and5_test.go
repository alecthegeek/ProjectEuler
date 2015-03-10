// Breif one line overview of package
package multiplesOf3and5

import "testing"
import "fmt"

var data = []struct{ n, r uint }{
	{0, 0},
	{1, 0},
	{3, 0},
	{4, 3},
	{10, 23},
}

func TestProblem1(t *testing.T) {

	for _, d := range data {
		if got := Problem1(d.n); got != d.r {
			t.Errorf(fmt.Sprintf("Problem1 of %d returned %d, expected %d\n", d.n, got, d.r))
		}
	}
}
