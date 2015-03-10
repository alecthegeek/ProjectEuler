// Brief one line overview of package
package EvenFibonacciNumbers

//import "fmt"

func Problem2(l uint) uint {

	if l < 2 {
		return 0
	}
	var (
		efs uint = 0
		cf  uint = 1 // last fib calcalated (fib(1)
		of  uint = 1 // Previous fib calculated (fibo(0)
	)

	for i := uint(2); cf < l; i++ {
		nf := of + cf
		if nf < l && nf%2 == 0 {
			//fmt.Printf("Found an even fib %d\n", nf)
			efs += nf
		}
		of, cf = cf, nf
	}

	return efs
}
