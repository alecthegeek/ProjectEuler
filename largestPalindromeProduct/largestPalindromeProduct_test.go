// Breif one line overview of package
package largestPalindromeProduct

import "testing"
import "fmt"

var problem4TestData = []struct {
	n, r1, r2, r3 uint
}{
	{2, 9009, 91, 99},
}

var palidroneTestData = []struct {
	s string
	r bool
}{
	{"9009", true},
	{"90109", true},
	{"9988", false},
	{"5776", false},
	{"9779", true},
	{"96799", false},
}

func TestIsPalidrone(t *testing.T) {

	for _, d := range palidroneTestData {
		if got := isPalindrone(d.s); got != d.r {
			t.Errorf(fmt.Sprintf("isPalidrone of %s returned %t, expected %t\n", d.s, got, d.r))
		}
	}
}

func TestProblem4(t *testing.T) {

	for _, d := range problem4TestData {
		if prod, f1, f2 := Problem4(d.n); prod != d.r1 {
			t.Errorf(fmt.Sprintf("Problem4 of %d returned %d (%dx%d), expected %d\n", d.n, prod, f1, f2, d.r1))
		}
	}
}
